package handler

import (
	"net/http"
	"log"

	"github.com/gin-gonic/gin"
	"book-server/internal/model"
)

var books = []model.Book{
    {ID: "1", Title: "Life of Pi", Author: "Yann Martel", PublicationYear: 2001, Isbn: "0-676-97376-0"},
    {ID: "2", Title: "The Kite Runner", Author: "Khaled Hosseini", PublicationYear: 2003, Isbn: "1-57322-245-3"},
    {ID: "3", Title: "The Pragmatic Programmer", Author: "Andrew Hunt and David Thomas", PublicationYear: 1999, Isbn: "978-0135957059"},
}

func RegisterRoutes(r *gin.Engine) {
	r.GET("/books", getBooks)
	r.GET("/books/:id", getBookByID)
	r.POST("/books", createBook)
}

func getBooks(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, books)
}

func getBookByID(c *gin.Context) {
    id := c.Param("id")

    for _, book := range books {
        if book.ID == id {
            c.IndentedJSON(http.StatusOK, book)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found!"})
}

func createBook(c *gin.Context) {
    var newBook model.Book

    if err := c.ShouldBindJSON(&newBook); err != nil {
        log.Println("Client error occurred. " + err.Error())
    	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error: " + err.Error()})
        return
    }

    for _, book := range books {
        if book.ID == newBook.ID {
            log.Printf("Book ID=%s already exists. New book creation failed.\n", newBook.ID)
            c.IndentedJSON(http.StatusConflict, gin.H{"message": "Book ID already exists!"})
            return
        }
    }

    books = append(books, newBook)
    log.Printf("New book created: ID=%s, Title=%s\n", newBook.ID, newBook.Title)
    c.IndentedJSON(http.StatusCreated, newBook)
}
