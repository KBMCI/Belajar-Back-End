package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"belajar-back-end/book"
)

// membuat function handler agar lebih ringkas
func RootHandler(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"title": "ini title",
		"content": "ini content",
	})
}

func HelloHandler(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

func BooksHandler(c *gin.Context){
	// untuk menangkap parameter path variable id
	id := c.Param("id")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H{
		"id": id,
		"title": title,
	})
}

func QueryHandler(c *gin.Context){
	// untuk menangkap parameter query string variable title
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"price": price,
	})
}

func PostBooksHandler(c *gin.Context)  {
	var bookInput book.BookInput

	err :=c.ShouldBindJSON(&bookInput)
	if err != nil {
		// untuk menampilkan error di response
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title": bookInput.Title,
		"price": bookInput.Price,
		"sub_title": bookInput.SubTitle,
	})
}