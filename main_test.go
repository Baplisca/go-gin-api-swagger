package main
import (
    "bytes"
    "testing"
    "encoding/json"
    "math"
    "net/http"
    "net/http/httptest"
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
)

// ref : https://circleci.com/blog/gin-gonic-testing/
func SetUpRouter() *gin.Engine{
    router := gin.Default()
    return router
}

func TestCreateBookHandler(t *testing.T) {
    r := SetUpRouter()
    r.POST("/books", createBook)
    book := book{
        ID: "100",
        Title: "It's my life",
        Author: "Baplisca",
        Quantity: math.MaxInt8,
    }
    jsonValue, _ := json.Marshal(book)
    req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(jsonValue))

    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusCreated, w.Code)
}