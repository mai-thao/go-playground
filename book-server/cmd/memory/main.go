package main

import (
    "github.com/gin-gonic/gin"
    "book-server/internal/handler"
    "book-server/internal/database"
)

func main() {
    database.ConnectStandard()
    database.ConnectOrm()
    defer database.Db.Close()

    router := gin.Default()
    handler.RegisterRoutes(router)
    router.Run("localhost:8080")
}
