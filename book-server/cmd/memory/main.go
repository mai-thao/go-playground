package main

import (
    "log"
    "database/sql"

    _ "github.com/lib/pq"
    "github.com/gin-gonic/gin"
    "book-server/internal/handler"
)

func main() {
    log.Println("Attempting to connect to Postgres database...")
    connStr := "postgres://adminmt:adminpw@localhost:5432/bookdb?sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    defer db.Close()
    if err != nil {
        log.Fatal(err)
    }
    if err = db.Ping(); err != nil {
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

    _, err = db.Exec(tableCreationSql)
    if err != nil {
        log.Println("Database migration failed")
        log.Fatal(err)
    }
    log.Println("Successfully created new books table!")

    router := gin.Default()
    handler.RegisterRoutes(router)
    router.Run("localhost:8080")
}
