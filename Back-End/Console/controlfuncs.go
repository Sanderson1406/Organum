package console

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func GetCat() string {
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
