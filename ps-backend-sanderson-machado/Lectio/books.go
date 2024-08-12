package lectio

type Book struct {
	idBook   string  `json:"idBook"  bson:"id"`
	Title    string  `json:"title"	bson:"title"`
	Author   string  `json:"author"	bson:"author"`
	Category string  `json:"category"	bson: "category"`
	Pages    float32 `json:"pages"	bson:"pages"`
}

type Review struct {
	idBook string
	opnion string
	score  float32
}
