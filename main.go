package main

import (
	_ "encoding/json"
	_ "fmt"
	"log"
	_ "net/http"

	"Belajar-Back-End/book"
	"Belajar-Back-End/handler"

	"github.com/gin-gonic/gin"
	_ "github.com/go-playground/validator/v10"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
) // anjir kenapa importnya harus manual ya sksksk :(

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/belajar-back-end?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB connection error")
	}
	// ======================================================================================
	// 										MIGRATE DATABASE
	// ======================================================================================
	db.AutoMigrate(&book.Book{}) 

	// ======================================================================================
	// 											 CRUD 
	// ======================================================================================


	// ======================================================================================
	// 											CREATE
	// ======================================================================================

	// Struct cara 2 
	// book  := book.Book{
	// 	Title: "Manusia setengah ayam",
	// 	Price: 50000,
	// 	Discount: 50,
	// 	Rating: 4,
	// 	Description: "Bukunya sangat jelek",
	// }

	// struct cara 1
	// book := book.Book{}
	
	// book.Title = "Manusia setengah salmon"
	// book.Price = 120000
	// book.Discount = 10
	// book.Rating = 5
	// book.Description = "Bukunya sangat bagus"

	// err = db.Create(&book).Error

	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error creating book record")
	// 	fmt.Println("==========================")
	// }

	// ======================================================================================
	// 										READ / QUERY 
	// ======================================================================================

	// Read data pertama 

	// var book book.Book

	// err = db.Debug().First(&book).Error 		// mendapatkan data pertama
	// err = db.Debug().Last(&book).Error 		// mendapatkan data terakhir
	// err = db.Debug().First(&book, 2).Error 	// mendapatkan data berdasaarkan primary key

	// 	if err != nil {
	// 			fmt.Println("==========================")
	// 			fmt.Println("Error finding book record")
	// 			fmt.Println("==========================")
	// 	}

	// fmt.Println("Title book :", book.Title)
	// fmt.Println("book object %v", book)

	// Pengambilan seluruh data berupa slice 
	
	// var books []book.Book
	// err = db.Debug().Find(&books).Error
	
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error finding book record")
	// 	fmt.Println("==========================")
	// }

	// for _, b := range books {
	// 	fmt.Println("Title book :", b.Title)
	// 	fmt.Println("Book object %v:", b)
	// }

	// Read berdasarkan titlenya
	// var books []book.Book
	// err = db.Debug().Where("title = ?", "Manusia setengah ikan").Error
	
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error finding book record")
	// 	fmt.Println("==========================")
	// }

	// for _, b := range books {
	// 	fmt.Println("Title book :", b.Title)
	// 	fmt.Println("Book object %v:", b)
	// }

	// Read berdasarkan ratingnya
	// var books []book.Book
	// err = db.Debug().Where("rating = ?", 5).Error
	
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error finding book record")
	// 	fmt.Println("==========================")
	// }

	// for _, b := range books {
	// 	fmt.Println("Title book :", b.Title)
	// 	fmt.Println("Book object %v:", b)
	// }

	// ======================================================================================
	//											UPDATE
	// ======================================================================================

	// var book book.Book

	// err = db.Debug().Where("id = ?", 1).First(&book).Error
	
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error finding book record")
	// 	fmt.Println("==========================")
	// }

	// book.Title = "Manusia setengah kambing"

	// err = db.Save(&book).Error

	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error updating book record")
	// 	fmt.Println("==========================")
	// }
	
	// ======================================================================================
	// 											DELETE 
	// ======================================================================================

	// var book book.Book

	// err = db.Debug().Where("id = ?", 1).First(&book).Error
	
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error finding book record")
	// 	fmt.Println("==========================")
	// }

	// err = db.Delete(&book).Error

	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error updating book record")
	// 	fmt.Println("==========================")
	// }

	// DOKUMENTASI LENGKAP CRUD ADA DI https://gorm.io/docs/query.html

	// go get -u github.com/gin-gonic/gin untuk import 
	
	// membuat route 
	route := gin.Default()
	
	// contoh route sedrhana tapi tidak umum 
	// route.GET("/", func(c *gin.Context){
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"nama": "Aldy Imam Wijaya",
	// 		"bio": "Pengangguran banyak acara",
	// 	})
	// })

	// ======================================================================================
	// 								MEMBUAT URL PARAMETER 
	// ======================================================================================
	// route GET
	
	// route.GET("/", rootHandler)
	// route.GET("/hello", helloHandler)
	// route.GET("/books/:id/:title", booksHandler) 
	// route.GET("/query", queryHandler)

	// ======================================================================================
	// 									REPOSITORY LAYER
	// ======================================================================================
	
	bookRepository := book.NewRepository(db)

	// books, err := bookRepository.FindAll()

	// for _, book := range books {
	// 	fmt.Println("Title :", book.Title)
	// }

	// book, err := bookRepository.FindByID(2)
	// fmt.Println("Title :", book.Title)

	// 	book := book.Book{
	// 		Title: "Manusia setengah T-Rex",
	// 		Price: 150000,
	// 		Rating: 5,
	// 		Description: "Buku ini biasa saja",
	// 		Discount: 0,
	// 	}
	// bookRepository.CreateBook(book)

	// ======================================================================================
	// 									SERVICE LAYER
	// ======================================================================================

	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)
	// bookRequest := book.BookRequest{
	// 	Title: "Manusia setengah anjing",
	// 	Price: "35000",
	// }

	// bookService.Create(bookRequest)

	// versioning / grouping
	v1 := route.Group("/v1")

	// v1.GET("/", bookHandler.RootHandler)
	// v1.GET("/hello", bookHandler.HelloHandler)
	// v1.GET("/books/:id/:title", bookHandler.BooksHandler) // tanda " : " merupakan sebuah variable
	// v1.GET("/query", bookHandler.QueryHandler)


	// route GET --> untuk mendapatkan
	v1.GET("/books", bookHandler.GetBooks)
	v1.GET("/books/:id", bookHandler.GetBook)

	// route POST
	v1.POST("/books", bookHandler.CreateBook)
	
	//route put --> untuk update
	v1.PUT("/books/:id", bookHandler.UpdateBook)
	
	// route DELETE 
	v1.DELETE("/books/:id", bookHandler.DeleteBook)

	route.Run() // mengubah port (":8888"), default port (8080)
}

