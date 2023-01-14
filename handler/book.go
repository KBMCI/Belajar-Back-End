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
func RootHandler(c *gin.Context) {
	// c.JSON digunakan untuk mengirim data
	c.JSON(http.StatusOK, gin.H{
		"name": "Wisnu",
		"bio":  "Aktor film meteor",
	})
}

func BooksHandler(c *gin.Context) {
	//untuk menangkap variable menggunakan param
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func QueryHandler(c *gin.Context) {
	//untuk menangkap query string menggunakan Query()
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"price": price,
	})
}

func PostBooksHandler(c *gin.Context) {
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

	c.JSON(http.StatusOK, gin.H{
		"title":     BookRequest.Title,
		"price":     BookRequest.Price,
		"sub_title": BookRequest.SubTitle,
	})
}