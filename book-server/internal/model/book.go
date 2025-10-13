package model

type Book struct {
    ID              int     `json:"id"`
    Title           string  `json:"title"`
    Author          string  `json:"author"`
    PublicationYear int     `json:"publication_year"`
    Isbn            string  `json:"isbn"`
}
