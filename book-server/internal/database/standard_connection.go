package database

import (
    "log"
    "database/sql"
    "math/rand"

    _ "github.com/lib/pq"
    "book-server/internal/model"
)

var Db *sql.DB

// ConnectStandard connects and persists data to Postgres DB using Go's built in standard SQL package (raw SQL)
func ConnectStandard() {
    log.Println("Attempting to connect to Postgres database...")
    connStr := "postgres://adminmt:adminpw@localhost:5432/bookdb?sslmode=disable"
    Db, err := sql.Open("postgres", connStr)

    if err != nil {
        log.Fatal(err)
    }
    if err = Db.Ping(); err != nil {
        log.Println("Database ping failed")
        log.Fatal(err)
    }
    log.Println("Successfully connected to Postgres database!")

    tableCreationSql := `
        CREATE TABLE IF NOT EXISTS books (
            id SERIAL PRIMARY KEY,
            title VARCHAR,
            author VARCHAR,
            publication_year INTEGER,
            isbn VARCHAR
        );
    `

    _, err = Db.Exec(tableCreationSql)
    if err != nil {
        log.Println("Database migration failed")
        log.Fatal(err)
    }
    log.Println("Successfully created new books table!")

    book := model.Book{
        ID: rand.Intn(100),
        Title: "The Pragmatic Programmer",
        Author: "Andrew Hunt and David Thomas",
        PublicationYear: 1999,
        Isbn: "978-0135957059",
    }

    insertSql := `
        INSERT INTO books (
            id, title, author, publication_year, isbn
        ) VALUES (
            $1, $2, $3, $4, $5
        ) RETURNING id
    `

    err = Db.QueryRow(insertSql, book.ID, book.Title, book.Author, book.PublicationYear, book.Isbn).Scan(&book.ID)
    if err != nil {
        log.Println("Database insertion failed")
        log.Fatal(err)
    }

    log.Printf("Inserted book with ID: %d", book.ID)
}
