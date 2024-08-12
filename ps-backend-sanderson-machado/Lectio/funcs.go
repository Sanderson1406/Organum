package main

import "github.com/google/uuid"

type Operations struct{}

func (a Operations) AddBook(title, author, category string, pages float32) (*Book, string) {
	code := uuid.New().String()
	newBook := Book{
		idBook:   code,
		Title:    title,
		Author:   author,
		Category: category,
		Pages:    pages,
	}
	
	return &newBook, "Livro adicionado com sucesso!"
}

func (b Operations) Update(book *Book, newStatus, nome string) (*Book, string) {
	id, _ := TitletoID(nome)
	for i, book := range books {
		if book.idBook == id {
			books[i].Category = newStatus
			return &books[i], "Atualizado com sucesso"
		}
	}
	return nil, "Livro não encontrado"
}

func (c Operations) List(books []Book, cate string) []string {
    toreturn := make([]string, 0)
    for i, book := range books {
        if book.Category == cate {
            toreturn = append(toreturn, books[i].Title)
        }
    }
    return toreturn
}

func (d Operations) Revi(b *Book, name, op string, sc float32) (string, *Review) {
    id, _ := TitletoID(name) 
    newReview := Review{
		idBook:	id,
		opnion: op,
		score: sc,
	}
    return "Review e nota adicionado com sucesso", &newReview
}

func TitletoID(title string) (string, string) {
	for _, book := range books {
		if book.Title == title {
			return book.idBook, "Livro encontrado!"
		}
	}
	return "", "Livro não encontrado"
}