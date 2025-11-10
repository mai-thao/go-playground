package handler

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
)

// Test for checking responsive webserver
func TestPingRoute(t *testing.T) {
    router := gin.New()
    RegisterRoutes(router)

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/ping", nil)
    router.ServeHTTP(w, req)

    if w.Code != http.StatusOK {
        t.Errorf("Expected status 200, got %d", w.Code)
    }

    if w.Body.String() != "pong" {
        t.Errorf("Expected body %q, got %q", "pong", w.Body.String())
    }
}

// // Test for GET /books
// func TestGetBooks(t *testing.T) {
//
// }
//
// // Test for GET /books/:id
// func TestGetBookByID(t *testing.T) {
//
// }
//
// // Test for POST /books
// func TestCreateBook(t *testing.T) {
//
// }
//
// // Test for DELETE /books/:id
// func TestDeleteBook(t *testing.T) {
//
// }
