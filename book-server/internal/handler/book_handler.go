package handler

import (
	"net/http"
	"log"
	"database/sql"

	"github.com/gin-gonic/gin"
	"book-server/internal/model"
    "book-server/internal/database"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/books", getBooks)
	r.GET("/books/:id", getBookByID)
	r.POST("/books", createBook)
}

// Helpful Go data querying doc: https://go.dev/doc/database/querying
func getBooks(c *gin.Context) {
    rows, err := database.Db.Query("SELECT * FROM books")
    if err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error: " + err.Error()})
        return
    }
    defer rows.Close()

    var books []model.Book
    for rows.Next() {
        var book model.Book
        if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.PublicationYear, &book.Isbn); err != nil {
            c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error: " + err.Error()})
            return
        }
        books = append(books, book)
    }
    log.Printf("Retrieved %d books\n", len(books))

    c.IndentedJSON(http.StatusOK, books)
}


func getBookByID(c *gin.Context) {
    var book model.Book
    id := c.Param("id")
    row := database.Db.QueryRow("SELECT * FROM books WHERE id = $1", id)
    err := row.Scan(&book)
    if err != nil {
        if err == sql.ErrNoRows {
            c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found!"})
            return
        }
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error: " + err.Error()})
        return
    }

    c.IndentedJSON(http.StatusOK, book)
}

func createBook(c *gin.Context) {
    var newBook model.Book

    if err := c.ShouldBindJSON(&newBook); err != nil {
        log.Println("Client error occurred. " + err.Error())
    	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error: " + err.Error()})
        return
    }

    insertSql := `
        INSERT INTO books (
            id, title, author, publication_year, isbn
        ) VALUES (
            $1, $2, $3, $4, $5
        ) RETURNING id
    `
    err := database.Db.QueryRow(insertSql, newBook.ID, newBook.Title, newBook.Author, newBook.PublicationYear, newBook.Isbn).Scan(&newBook.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Error: " + err.Error()})
        return
    }
    log.Printf("New book created: ID=%d, Title=%s\n", newBook.ID, newBook.Title)
    c.IndentedJSON(http.StatusCreated, newBook)
}
