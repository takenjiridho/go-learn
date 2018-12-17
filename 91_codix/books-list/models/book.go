package models

type Rdata struct {
	Status string `json:"status"`
	Data   []Book `json:"data"`
}

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

// type v struct {
// 	Books []Book `json:"data"`
// }
