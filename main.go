package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
    swaggerFiles "github.com/swaggo/files"
    "github.com/swaggo/gin-swagger"
	_ "example/go-simple-api/docs"
)

type book struct{
	ID	string	`json:"id"`
	Title	string	`json:"title"`
	Author	string	`json:"author"`
	Quantity	int	`json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}


// @Schemes
// @Description return all Books
// @Tags books
// @Accept json
// @Produce json
// @Router /books [get]
// @Success 200
func getBooks(c *gin.Context){
	c.IndentedJSON(http.StatusOK, books)
}


// @Schemes
// @Description return specific book
// @Tags books
// @Accept json
// @Produce json
// @Router /books/:id [get]
// @Success 200
// @Failure 404
func bookById(c *gin.Context){
	id := c.Param("id")
	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book is not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}


// @Schemes
// @Description register new book
// @Tags books
// @Accept json
// @Produce json
// @Router /books [post]
// @Success 201
func createBook(c *gin.Context){
	var newBook book 
	if err := c.BindJSON(&newBook); err != nil{
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}


// @Schemes
// @Description rent a specific book
// @Tags checkout
// @Accept json
// @Produce json
// @Router /checkout [patch]
// @Success 201
// @Failure 400
// @Failure 404
func checkoutBook(c *gin.Context){
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"missing id query parameter"})
		return
	}

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book is not found"})
		return
	}

	if book.Quantity <= 0{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"book is not available"})
		return
	}
	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

// @Schemes
// @Description return a specific book
// @Tags return
// @Accept json
// @Produce json
// @Router /return [patch]
// @Success 201
// @Failure 400
// @Failure 404
func returnBook(c *gin.Context){
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"missing id query parameter"})
		return
	}

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book is not found"})
		return
	}

	book.Quantity++
	c.IndentedJSON(http.StatusOK, book)
}

func getBookById(id string) (*book, error){
	for i,b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("book is not found")
}

func main(){
	router := gin.Default()

	// swagger setting
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)
	router.POST("/books", createBook)
	router.PATCH("/checkout", checkoutBook)
	router.PATCH("/return", returnBook)
	router.Run("localhost:8080")
}