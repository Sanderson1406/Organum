package models

import "github.com/dgrijalva/jwt-go"

type GetUser struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type User struct {
	ID        string `json:"id" bson:"_id"`
	Email     string `json:"email" bson:"email"`
	Name      string `json:"name" bson:"name"`
	Sobrenome string `json:"sobrenome" bson:"sobrenome"`
	Password  string `json:"password" bson:"password"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
