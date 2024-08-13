package main

import (
	"fmt"
	"log"
	"time"
	"net/http"
	"strconv"
	//"strings"
	//"encoding/json"
	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
	"context"
)

var bookss []Book
var reviewss []Review
var client *mongo.Client

type myHandler struct{}

func (h myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && r.URL.Path == "/adicionarbook" {
		op := "Adicionar"
		title := r.URL.Query().Get("title")
		autor := r.URL.Query().Get("author")
		cat := r.URL.Query().Get("status")
		pg := r.URL.Query().Get("pags")

		if title == "" || autor == "" || cat == "" || pg == "" {
			http.Error(w, "Parâmetros inválidos", http.StatusBadRequest)
			return
		}

		f, err := strconv.ParseFloat(pg, 32)
		if err != nil {
			http.Error(w, "Erro ao converter páginas para número: "+err.Error(), http.StatusBadRequest)
			return
		}
		pgs := float32(f)

		var eu *Book
		eu, _ = opcoesbook[op].AddBook(title, autor, cat, pgs) 
	
		
		collection := client.Database("Organum").Collection("books")
		insertResult, err := collection.InsertOne(context.Background(), eu)
		if err != nil {
			http.Error(w, "Erro ao inserir livro no MongoDB"+err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "Livro inserido ", insertResult.InsertedID)
	}
	if r.Method == "GET" && r.URL.Path == "/listarbooks"  {
		for _, livro := range bookss {
			detalhesLivro := fmt.Sprintf("[ID: %s, Título: %s, Autor: %s, Páginas: %d\n, Categoria: %s]", livro.idBook, livro.Title, livro.Author, livro.Pages, livro.Category)
			w.Write([]byte(detalhesLivro))
		}
	}
}

	
	

func StartServer() {
	s := &http.Server{
		Addr:           "192.168.137.1:8080",
		Handler:        myHandler{},
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	log.Fatal(s.ListenAndServe())
}
