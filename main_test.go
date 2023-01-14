package main

import (
	"bytes"
	"encoding/json"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// ref : https://circleci.com/blog/gin-gonic-testing/
func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetBooks(t *testing.T) {
	// set up
	r := SetUpRouter()
	r.GET("/books", getBooks)

	// execute
	req, _ := http.NewRequest("GET", "/books", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// assert
	var books []book
	json.Unmarshal(w.Body.Bytes(), &books)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, books)
}

func TestCreateBook(t *testing.T) {
	// set up
	r := SetUpRouter()
	r.POST("/books", createBook)

	// mock
	book := book{
		ID:       "100",
		Title:    "It's my life",
		Author:   "Baplisca",
		Quantity: math.MaxInt8,
	}
	jsonValue, _ := json.Marshal(book)

	// execute
	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// assert
	assert.Equal(t, http.StatusCreated, w.Code)
}
