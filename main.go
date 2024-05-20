package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)


type book struct{
	ID string	`json:"id"`
	Title string	`json:"title"`
	Author string	`json:"author"`
	Quantity int	`json:"quntity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

func getBooks(c *gin.Context){
	c.IndentedJSON(http.StatusOK, books)
}

func createBook(c *gin.Context){
	var newBook book
	//Here we are going to bind the data which is going to be received in the payload 
	//of the POST request to the json in our code
	if err := c.BindJSON(&newBook);err != nil{
		return 
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)

}

func getBookbyId(id string) (*book, error){
	for i, b := range books{
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found!")

}

func bookById(c *gin.Context){
	id := c.Param("id")
	books, err := getBookbyId(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}
	c.IndentedJSON(http.StatusOK, books)
}

func checkoutBook( c *gin.Context){
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"Missing id query parameter."})
		return
	}
	book, err := getBookbyId(id)

	if err != nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "Book not found."})
		return
	}
	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "Book not available."})
		return
	}
	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

func returnBooks( c *gin.Context){
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"Missing id query parameter."})
		return
	}
	book, err := getBookbyId(id)

	if err != nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "Book not found."})
		return
	}
	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)


}

func main(){
	// we are generating a GET endpoint here
	router := gin.Default()
	//used command: curl localhost:8080/books
	router.GET("/books", getBooks)
	//used command: curl localhost:8080/books --include --header "Content-Type: application/json" -d @body.json --request "POST"
	//d: data, we provide the data with "@" followed by the file
	router.GET("/books/:id", bookById)
	router.POST("/books", createBook)
	//used command: curl -s -X PATCH 'localhost:8080/checkout?id=2'
	router.PATCH("/checkout", checkoutBook)
	router.PATCH("/return", returnBooks)
	router.Run("localhost:8080")

}