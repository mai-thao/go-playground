package database

import (
    "log"
    "math/rand"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "book-server/internal/model"
)

// ConnectOrm connects and persists data to Postgres DB using GORM as the Object-Relational Mapper layer
func ConnectOrm() {
    connStr := "postgres://adminmt:adminpw@localhost:5432/bookdb?sslmode=disable"
    db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

    if err != nil {
        log.Println("Database connection via ORM failed")
        log.Fatal(err)
    }

    book := model.Book{
        ID: rand.Intn(100),
        Title: "The Pragmatic Programmer",
        Author: "Andrew Hunt and David Thomas",
        PublicationYear: 1999,
        Isbn: "978-0135957059",
    }
    db.Create(&book)
    log.Printf("Inserted book via ORM with ID: %d", book.ID)

    var books []model.Book
    db.Find(&books)

    log.Printf("Queried for %d books via ORM", len(books))
}
