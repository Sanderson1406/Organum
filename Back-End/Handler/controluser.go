package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	database "organum/DataBase"
	models "organum/Models"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var SECRET_KEY = []byte("secret")
var BlacklistedTokens = make(map[string]bool)

func GenerateJWT(email string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &models.Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	user.ID = models.GenerateID()

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Erro ao receber dados: %v", err)
		return
	}

	_, err = CreateUserDB(user)
	if err != nil {
		if err.Error() == "o e-mail já está em uso" {
			http.Error(w, err.Error(), http.StatusConflict)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		log.Printf("Erro ao criar usuário: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Erro ao codificar a resposta: %v", err)
		return
	}
}

func CreateUserDB(user models.User) (*mongo.InsertOneResult, error) {
	client := database.GetMongoClient()
	collection := client.Database("Organum").Collection("profiles")
	filter := bson.M{"email": user.Email}
	var existingUser models.User

	err := collection.FindOne(context.Background(), filter).Decode(&existingUser)
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, err
	}

	if existingUser.Email != "" {
		return nil, errors.New("o e-mail já está em uso")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.ID = models.GenerateID()
	user.Password = string(hashedPassword)

	insertResult, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}

	return insertResult, nil
}

func Login(w http.ResponseWriter, r *http.Request) {
	var gu models.GetUser
	err := json.NewDecoder(r.Body).Decode(&gu)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Erro ao decodificar a solicitação: %v", err)
		return
	}

	exists, err := LoginDB(gu)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Erro ao verificar o usuário: %v", err)
		return
	}

	if exists {

		tokenString, err := GenerateJWT(gu.Email)
		if err != nil {
			http.Error(w, "Erro ao gerar token", http.StatusInternalServerError)
			log.Printf("Erro ao gerar token: %v", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
		log.Printf("Login bem-sucedido para o usuário: %v", gu.Email)
	} else {
		http.Error(w, "Usuário ou senha inválidos", http.StatusUnauthorized)
	}
}

func LoginDB(gu models.GetUser) (bool, error) {
	client := database.GetMongoClient()
	collection := client.Database("Organum").Collection("profiles")
	filter := bson.M{"email": gu.Email}
	var user models.User

	err := collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("Usuário não encontrado com o email fornecido.")
			return false, nil
		}
		log.Printf("Erro ao buscar usuário: %v", err)
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(gu.Password))
	if err != nil {
		log.Printf("Erro ao comparar a senha: %v", err)
		return false, fmt.Errorf("senha incorreta")
	}

	return true, nil
}

func Logout(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Cabeçalho de Autorização é necessário", http.StatusUnauthorized)
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		http.Error(w, "Token Bearer é necessário", http.StatusUnauthorized)
		return
	}

	BlacklistedTokens[tokenString] = true

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Logout realizado com sucesso"))
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Cabeçalho de Autorização é necessário", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			http.Error(w, "Token Bearer é necessário", http.StatusUnauthorized)
			return
		}

		if BlacklistedTokens[tokenString] {
			http.Error(w, "Token está na lista negra (usuário deslogado)", http.StatusUnauthorized)
			return
		}

		claims := &models.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return SECRET_KEY, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				http.Error(w, "Assinatura do token inválida", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}
		if !token.Valid {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "userEmail", claims.Email)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
