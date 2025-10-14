package database

import (
    "log"
    "math/rand"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "book-server/internal/model"
)

func ConnectOrm() {
    connStr := "postgres://adminmt:adminpw@localhost:5432/bookdb?sslmode=disable"
    db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

    if err != nil {
        log.Println("Database connection via ORM failed")
        log.Fatal(err)
    }

    db.AutoMigrate(&model.Book{})
    book := model.Book{
        ID: rand.Intn(100),
        Title: "The Pragmatic Programmer",
        Author: "Andrew Hunt and David Thomas",
        PublicationYear: 1999,
        Isbn: "978-0135957059",
    }
    db.Create(&book)
    log.Printf("Inserted book via ORM with ID: %d", book.ID)
}
