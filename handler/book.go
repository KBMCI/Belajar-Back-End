package handler

import (
	_ "encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"Belajar-Back-End/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
) // anjir kenapa importnya harus manual ya sksksk :(

// handler

type bookHandler struct {
	bookService book.Service	
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) GetBooks(c *gin.Context) {
	books, err := h.bookService.FindAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
			})
			return
	}

	var booksResponse []book.BookResponse

	for _, b := range books {
		bookResponse := convertToBookResponse(b)

		booksResponse = append(booksResponse, bookResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})
}

func (h *bookHandler) GetBook(c *gin.Context)  {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	b, err := h.bookService.FindByID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
			})
			return
	}

	bookResponse := convertToBookResponse(b)

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

func (h *bookHandler) DeleteBook(c *gin.Context)  {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	b, err := h.bookService.Delete(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
			})
			return
	}

	bookResponse := convertToBookResponse(b)

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

// func (handler *bookHandler) RootHandler(c *gin.Context)  {
// 	c.JSON(http.StatusOK, gin.H{
// 		"nama": "Aldy Imam Wijaya",
// 		"bio": "Pengangguran banyak acara",
// 	})
// }

// func (h *bookHandler) HelloHandler(c *gin.Context)  {
// 	c.JSON(http.StatusOK, gin.H{
// 		"content": "Hello World",
// 		"subtitle": "Belajar Golang",
// 	}) 
// }

// func (h *bookHandler) BooksHandler(c *gin.Context)  {
// 	id := c.Param("id") // untuk mengambil path variable menggunakan Param
// 	title := c.Param("title")

// 	c.JSON(http.StatusOK, gin.H{
// 		"id": id,
// 		"title": title,
// 	})
// }

// func (h *bookHandler) QueryHandler(c *gin.Context)  {
// 	title := c.Query("title") // untuk mengambil query string
// 	price := c.Query("price")
// 	c.JSON(http.StatusOK, gin.H{
// 		"title": title,
// 		"price": price,
// 	})
// }

func (h *bookHandler) CreateBook(c *gin.Context){
	
	// title, price
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)

	if err != nil {
		// log.Fatal(err) // akan menghentikan server

		// menampilkan 1 error
		// for _, v := range err.(validator.ValidationErrors) {
		// 	errorMassage := fmt.Sprintf("Error on field %s, condition: %s", v.Field(), v.ActualTag())
		// 	c.JSON(http.StatusBadRequest, errorMassage)
		// 	return
		// }

		// menampilkan lebih dari 1 error
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
			})
			return
	}

	book, err := h.bookService.Create(bookRequest)


	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
			})
			return
	}
		c.JSON(http.StatusOK, gin.H{
			"data": convertToBookResponse(book),
			// "price": bookInput.Price,
			// "sub_title": bookInput.SubTitle,
		})
	
}

func (h *bookHandler) UpdateBook(c *gin.Context){
	
	// title, price
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)

	if err != nil {

		// menampilkan lebih dari 1 error
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
			})
			return
	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	book, err := h.bookService.Update(id, bookRequest)


	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
			})
			return
	}
		c.JSON(http.StatusOK, gin.H{
			"data": convertToBookResponse(book),
			// "price": bookInput.Price,
			// "sub_title": bookInput.SubTitle,
		})
	
}

func convertToBookResponse(b book.Book) book.BookResponse{
	return book.BookResponse{
		ID: 			b.ID,
		Title: 			b.Title,
		Description: 	b.Description,
		Price: 			b.Price,
		Rating: 		b.Rating,
		Discount: 		b.Discount,
	}
}