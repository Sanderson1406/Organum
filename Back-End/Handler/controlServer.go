package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"

	database "organum/DataBase"
	models "organum/Models"
)

func CreateAndAddNewBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Erro ao decodificar a solicitação: %v", err)
		return
	}

	book.ID = models.GenerateID()

	result, err := database.AddBookDB(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Erro ao adicionar livro ao banco de dados: %v", err)
		return
	}

	book.ID = result.InsertedID.(primitive.ObjectID).Hex()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Erro ao codificar a resposta: %v", err)
		return
	}
}

func ListBooksByStatus(w http.ResponseWriter, r *http.Request) {
	tag := r.URL.Query().Get("tag")

	selected, err := database.ListBooksByStatusBD(tag)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Erro ao listar livros por status: %v", err)
		return
	}

	var books []models.Book
	for _, t := range selected {
		book := models.Book{
			Title:  t.Title,
			Author: t.Author,
			Tag:    t.Tag,
			Pages:  t.Pages,
		}
		books = append(books, book)
	}

	err = json.NewEncoder(w).Encode(books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Erro ao codificar a resposta: %v", err)
		return
	}

}

func UpdateBookTag(w http.ResponseWriter, r *http.Request) {
	var booktag models.NewTag

	err := json.NewDecoder(r.Body).Decode(&booktag)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Erro ao decodificar a solicitação: %v", err)
		return
	}

	bookID, err := database.GetBookById(booktag.Title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Erro ao capturar ID: %v", err)
		return
	}

	result, err := database.UpdateBookTag(bookID, booktag.NewTag)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Erro ao atualizar tag: %v", err)
		return
	}

	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Erro ao codificar a resposta: %v", err)
		return
	}
}

func CreateReview(w http.ResponseWriter, r *http.Request) {
	var addreview models.AddReview

	err := json.NewDecoder(r.Body).Decode(&addreview)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Erro ao decodificar a solicitação: %v", err)
		return
	}

	newid := models.GenerateID()
	insertResult, err := database.CreateReviewDB(addreview, newid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Erro ao adicionar review ao banco de dados: %v", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Revisão criada com sucesso. ID: %s", insertResult.InsertedID)

}

func GetReview(w http.ResponseWriter, r *http.Request) {

	var getreview []models.AddReview

	getreview, err := database.GetReviewBD()

	err = json.NewEncoder(w).Encode(getreview)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Erro ao codificar a resposta: %v", err)
		return
	}

}

func UpdateBookProgress(w http.ResponseWriter, r *http.Request) {
	var putprogress models.PutProgress

	err := json.NewDecoder(r.Body).Decode(&putprogress)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Erro ao codificar")
		return
	}

	bookid, err := database.GetBookById(putprogress.Title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Erro ao capturar ID: %v", err)
		return
	}

	resutl, err := database.UpdateBookProgressDB(bookid, putprogress.Pgl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Erro ao atualizar progresso: %v", err)
		return
	}

	fmt.Fprintf(w, "Atualização feita com sucesso: %v\n", resutl)
	w.WriteHeader(http.StatusCreated)

}

func GetBookProgress(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")

	idbook, err := database.GetBookById(title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Erro ao capturar ID")
	}

	pags, err := database.GetReadPages(idbook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Erro ao capturar paginas")
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pags)
	fmt.Println("Enviado")

}

func CreateQuote(w http.ResponseWriter, r *http.Request) {
	var quote models.AddQuote

	err := json.NewDecoder(r.Body).Decode(&quote)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Erro ao decodificar a solicitação: %v", err)
		return
	}

	bookID, err := database.GetBookById(quote.Title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Erro ao capturar ID: %v", err)
		return
	}

	result, err := database.CreateQuoteBD(bookID, quote.Quote)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Erro ao atualizar tag: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func GetQuotes(w http.ResponseWriter, r *http.Request) {
	getquote, err := database.GetQuotesBD()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err != nil {
		log.Fatal(err)
	}
	err = json.NewEncoder(w).Encode(getquote)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Erro ao codificar a resposta: %v", err)
		return
	}
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	var title models.Title

	err := json.NewDecoder(r.Body).Decode(&title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Erro ao decodificar a solicitação: %v", err)
		return
	}

	idbook, err := database.GetBookById(title.Title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Erro ao capturar ID: %v", err)
		return
	}

	result1, err := database.DeleteBookDB(idbook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Erro ao deletar livro: %v", err)
		return
	}
	fmt.Fprintf(w, "Atualização feita com sucesso: %v\n", result1)
	w.WriteHeader(http.StatusCreated)

	result2, err := database.DeleteProgressDB(idbook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Erro ao deletar livro: %v", err)
		return
	}
	fmt.Fprintf(w, "Atualização feita com sucesso: %v\n", result2)
	w.WriteHeader(http.StatusCreated)

	result3, err := database.DeleteQuotesDB(idbook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Erro ao deletar livro: %v", err)
		return
	}
	fmt.Fprintf(w, "Atualização feita com sucesso: %v\n", result3)
	w.WriteHeader(http.StatusCreated)

	result4, err := database.DeleteReviewsDB(idbook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Erro ao deletar livro: %v", err)
		return
	}
	fmt.Fprintf(w, "Atualização feita com sucesso: %v\n", result4)
	w.WriteHeader(http.StatusCreated)
}
