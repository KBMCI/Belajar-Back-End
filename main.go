package main

import (
	"Belajar-Back-End/book"
	"Belajar-Back-End/handler"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Tambah framework gin dan gorm dengan mengetikkan
	// go get -u github.com/gin-gonic/gin
	// go get -u gorm.io/gorm

	// koneksi ke database
	dsn := "root:@tcp(127.0.0.1:3306)/belajar-back-end?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB connection error")
	}

	// auto migrate tabel
	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)

	bookRequest := book.BookRequest{
		Title:       "Gundam",
		Price:       "200000",
	}
	
	bookService.Create(bookRequest)


	// CRUD
	// CREATE
	/*
	book := book.Book{
		Title:       "Atomic habit",
		Description: "habit habit habit",
		Price:       120000,
		Rating:      5,
	}

	// db create digunakan untuk menambah data
	err = db.Create(&book).Error
	if err != nil {
		fmt.Println("Error creating book record")
	}

	// READ
	// var books []book.Book
	var book book.Book // digunakan untuk menyimpan data yang diambil dari db
	// err = db.First(&book).Error
	err = db.Find(&book).Error // menampilkan semua data
	// err = db.debug().First(&book).Error -> dapat memunculkan query yang digunakan
	if err != nil {
		fmt.Println("Error finding book record")
	}
	
	// for _, b := range books {
		fmt.Println("Title", book.Title)
		fmt.Println("book object", book)
	// }

	// UPDATE
	book.Title = "Man Tiger (Revised Edition)"
	err = db.Save(&book).Error
	if err != nil {
		fmt.Println("Error updating book record")
	}

	// DELETE
	err = db.Delete(&book).Error
	if err != nil {
		fmt.Println("Error deleting book record")
	}
	**/
	// membuat variable untuk menggunakan gin
	router := gin.Default()

	// membuat route
	/*
		router.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"name": "Wisnu",
				"bio":  "Aktor film meteor",
			})
		})
		**/

	// Method GET
	/*
		router.GET("/", rootHandler)
		router.GET("/books/:id", booksHandler) // :id -> sebuah variable, untuk menandakan variable menggunakan ":"
		router.GET("/query", queryHandler)

		// // Method POST
		router.POST("/books", postBooksHandler)
		**/

	// Versioning
	// Dapat digunakan apabila terjadi perubahan nama attribut pada json
	v1 := router.Group("/v1")
	v1.GET("/", handler.RootHandler)
	v1.GET("/books/:id", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)

	// menjalankan server
	router.Run() // -> port default 8080
	// router.Run(":8888") -> parameter untuk mendefenisikan port yang digunakan

}
