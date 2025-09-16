package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type book struct {
    ID              string  `json:"id"`
    Title           string  `json:"title"`
    Author          string  `json:"author"`
    PublicationYear int     `json:"publication_year"`
    Isbn            string  `json:"isbn"`
}

var books = []book{
    {ID: "1", Title: "Life of Pi", Author: "Yann Martel", PublicationYear: 2001, Isbn: "0-676-97376-0"},
    {ID: "2", Title: "The Kite Runner", Author: "Khaled Hosseini", PublicationYear: 2003, Isbn: "1-57322-245-3"},
    {ID: "3", Title: "The Pragmatic Programmer", Author: "Andrew Hunt and David Thomas", PublicationYear: 1999, Isbn: "978-0135957059"},
}

func getBooks(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, books)
}

func main() {
    router := gin.Default()
    router.GET("/books", getBooks)
    router.Run("localhost:8080")
}
