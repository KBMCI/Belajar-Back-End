package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"Belajar-Back-End/book"
)

// Membuat handler
// nama fungsi disesuaikan dengan nama route jika /hello nama : helloHandler

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) RootHandler(c *gin.Context) {
	// c.JSON digunakan untuk mengirim data
	c.JSON(http.StatusOK, gin.H{
		"name": "Wisnu",
		"bio":  "Aktor film meteor",
	})
}

func (h *bookHandler) BooksHandler(c *gin.Context) {
	//untuk menangkap variable menggunakan param
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (h *bookHandler) QueryHandler(c *gin.Context) {
	//untuk menangkap query string menggunakan Query()
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"price": price,
	})
}

func (h *bookHandler) PostBooksHandler(c *gin.Context) {
	var BookRequest book.BookRequest

	//Memasukkan data post ke bookInput
	err := c.ShouldBindJSON(&BookRequest) // menggunakan pointer karena kalo tidak maka bookInput yang menerima data hasil post hanya secara local variable
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		// dalam penggunaan c.json parameter pertama merupakan http code
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return

		// log.Fatal(err) -> log.fatal dapat menampilkan error yang disertai mengakhiri program
	}

	fmt.Println(BookRequest)

	book, err := h.bookService.Create(BookRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}
