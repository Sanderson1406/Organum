package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var books []Book
var reviews []Review

var opcoesbook = map[string]acs{
	"Adicionar": Operations{},
	"Atualizar": Operations{},
	"Listar":    nil,
	"Listar-Status": Operations{},
	"Adicionar-Review": Operations{},
	"Listar-Reviews": nil,
	//"":add{}
}

func Run() {
	for {
		var title, autor, cat, op, s string
		var pgs float32
		var eu *Book
		
		fmt.Println("Opções: ")
		for key := range opcoesbook {
			fmt.Println(key)
		}
		fmt.Scanln(&op)

		if op == "Adicionar" {
			title, autor, cat, pgs = Newbook()
			eu, s := opcoesbook[op].AddBook(title, autor, cat, pgs)
			books = append(books, *eu)
			fmt.Printf("ID do livro: %s\n", eu.idBook)
			fmt.Println(s)
		}
		if op == "Listar" {
			fmt.Println(books)
		}
		if op == "Atualizar" {
			newcatego, tochange := NewCat()
			_, s = opcoesbook[op].Update(eu, newcatego, tochange)
			fmt.Println(s)
			fmt.Println(books)
		}
		if op == "Listar-Status" {
			tolist := GetCat()
			cates := opcoesbook[op].List(books, tolist)
			fmt.Println(cates)
		}
		if op == "Adicionar-Review" {
			name, opn, sc := WriteReview()
			s, nr := opcoesbook[op].Revi(eu, name, opn, sc) 
			reviews = append(reviews, *nr) 
			fmt.Println(s)
			fmt.Println(nr)
		}
		if op == "Listar-Reviews" {
			fmt.Println(reviews)
		}
	}
}

func Newbook() (string, string, string, float32) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Titulo:")
	a, _ := reader.ReadString('\n')

	fmt.Println("Autor:")
	b, _ := reader.ReadString('\n')

	fmt.Println("Categoria:")
	c, _ := reader.ReadString('\n')

	fmt.Println("Paginas:")
	iStr, _ := reader.ReadString('\n')
	i, _ := strconv.ParseFloat(strings.TrimSpace(iStr), 32)
	return strings.TrimSpace(a), strings.TrimSpace(b), strings.TrimSpace(c), float32(i)
}

func NewCat() (string, string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Nova Categoria:")
	a, _ := reader.ReadString('\n')
	fmt.Println("Livro que deseja mudar:")
	b, _ := reader.ReadString('\n')
	return strings.TrimSpace(a), strings.TrimSpace(b) 
}

func GetCat() (string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Diga a categoria a ser listada:")
	a, _ := reader.ReadString('\n')
	return strings.TrimSpace(a)
}

func WriteReview() (string, string, float32) {
	reader := bufio.NewReader(os.Stdin)
	var c float32
	fmt.Println("Livro que deseja avaliar:")
	a, _ := reader.ReadString('\n')
	fmt.Println("Escreva sua opniao sobre o livro:")
	b, _ := reader.ReadString('\n')
	fmt.Println("Escreva a nota do livro (0-5):")
	fmt.Scanln(&c)
	return strings.TrimSpace(a), strings.TrimSpace(b), c
}