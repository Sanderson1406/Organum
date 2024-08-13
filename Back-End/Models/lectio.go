package models

type Review struct {
	ID     string  `json: "id"`
	Opnion string  `json: "opnion"`
	Score  float32 `json: "score"`
	IDbook string  `json: "idbook"`
}

type AddReview struct {
	Title  string  `json: "title"`
	Opnion string  `json: "opnion"`
	Score  float32 `json: "score"`
}

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Tag    string `json:"tag"`
	Pages  int    `json:"pages"`
	ID     string `json:"idBook"`
}

type NewTag struct {
	Title  string `json: "title"`
	NewTag string `json: "newtag"`
}

type Quote struct {
	IDBook string   `json: "idbook"`
	Quotes []string `json: "quote"`
}

type GetQuote struct {
	Title  string   `json: "title"`
	Quotes []string `json: "quote"`
}

type AddQuote struct {
	Title string `json: "title"`
	Quote string `json: "quote"`
}

type Title struct {
	Title string `json: "title"`
}

type Progress struct {
	IDbook string `json: "idbook"`
	ID     string `json: "id"`
	Pgt    int    `json: "pgt"`
	Pgl    int    `json: "pgl"`
}

type PutProgress struct {
	Title string `json: "title"`
	Pgl   int    `json: "pgl"`
}
