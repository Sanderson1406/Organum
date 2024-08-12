package lectio

type acs interface {
	AddBook(string, string, string, float32) (*Book, string)
	Update(*Book, string, string) (*Book, string)
	List([]Book, string) []string
	Revi(*Book, string, string, float32) (string, *Review)
}
