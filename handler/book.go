package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"belajar-back-end/book"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) RootHandler(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"title": "ini title",
		"content": "ini content",
	})
}

func (h *bookHandler) HelloHandler(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

func (h *bookHandler) BooksHandler(c *gin.Context){
	// untuk menangkap parameter path variable id
	id := c.Param("id")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H{
		"id": id,
		"title": title,
	})
}

func (h *bookHandler) QueryHandler(c *gin.Context){
	// untuk menangkap parameter query string variable title
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"price": price,
	})
}

// READ ALL BOOK
func (h *bookHandler) GetBooks(c *gin.Context) {
	books, err := h.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		
		return
	}

	// customize response
	var booksResponse []book.BookResponse

	for _, b := range books {
		bookResponse := convertToBookResponse(b)

		// memasukkan data ke slice booksResponse
		booksResponse = append(booksResponse, bookResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})
}

// READ BOOK BY ID
func (h *bookHandler) GetBook(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	b, err := h.bookService.FindById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		
		return
	}

	bookResponse := convertToBookResponse(b)

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})

}

// CREATE BOOK
func (h *bookHandler) PostBook(c *gin.Context)  {
	var bookRequest book.BookRequest

	err :=c.ShouldBindJSON(&bookRequest)
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

	book, err := h.bookService.CreateBook(bookRequest)

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

// UPDATE BOOK
func (h *bookHandler) UpdateBook(c *gin.Context)  {
	var updBookRequest book.UpdateBookRequest

	err :=c.ShouldBindJSON(&updBookRequest)
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

	// ambil parameter dulu di postman
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	book, err := h.bookService.UpdateBook(id, updBookRequest)

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

// DeLETE BOOK BY ID
func (h *bookHandler) DeleteBook(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	b, err := h.bookService.DeleteBook(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		
		return
	}

	bookResponse := convertToBookResponse(b)

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})

}

func convertToBookResponse(b book.Book) book.BookResponse  {
	return book.BookResponse{
		ID: b.ID,
		Title: b.Title,
		Description: b.Description,
		Price: b.Price,
		Rating: b.Rating,
		Discount: b.Discount,
	}
}