package main

import (
    "github.com/gin-gonic/gin"
    "book-server/internal/handler"
)

func main() {
    router := gin.Default()
    handler.RegisterRoutes(router)
    router.Run("localhost:8080")
}
