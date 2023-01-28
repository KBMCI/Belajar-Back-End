package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"belajar-back-end/book"
	"belajar-back-end/handler"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// connect database
	dsn := "root:@tcp(127.0.0.1:3306)/belajar-back-end?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	if err != nil {
		log.Fatal("Db connection error")
	}

	fmt.Println("Database connection success")
	db.AutoMigrate(&book.Book{})

	// ==============================
	// CRUD with layering repository
	// ==============================

	bookRepository := book.NewRepository(db)

	books, err := bookRepository.FindAll()
	if err != nil {
		fmt.Println("================================")
		fmt.Println("Error saat menampilkan data buku")
		fmt.Println("================================")
	}

	for _, book := range books {
		fmt.Println("Title :", book.Title)
	}

	idBook, err := bookRepository.FindById(1)
	if err != nil {
		fmt.Println("================================")
		fmt.Println("Error saat menampilkan data buku")
		fmt.Println("================================")
	}

	fmt.Println("Title :", idBook.Title)

	book := book.Book{
		Title: "Buku 3",
		Description: "Deskripsi buku 3",
		Price: 95000,
		Rating: 4,
		Discount: 0,
	}

	err = bookRepository.Create(book)
	if err != nil {
		fmt.Println("===============================")
		fmt.Println("Error saat memasukkan data buku")
		fmt.Println("===============================")
	}

	// ============================
	// CRUD not layering repository
	// ============================

	// ==========
	// CREATE
	// ==========
	// book := book.Book{}
	// book.Title = "Buku 2"
	// book.Price = 70000
	// book.Description = "Deskripsi buku 2"
	// book.Rating = 5
	// book.Discount = 10

	// err = db.Create(&book).Error
	// if err != nil {
	// 	fmt.Println("===============================")
	// 	fmt.Println("Error saat memasukkan data buku")
	// 	fmt.Println("===============================")
	// }
	
	// ==========
	// READ
	// ==========
	// var book book.Book

	// mengambil data buku pertama, Debug() untuk mengetahui query yg dijalankan
	// err = db.Debug().First(&book).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error saat read data buku")
	// 	fmt.Println("==========================")
	// }

	// fmt.Println("Title :", book.Title)
	// fmt.Println("book object : %v", book)

	// mengambil data buku terakhir
	// err = db.Debug().Last(&book).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error saat read data buku")
	// 	fmt.Println("==========================")
	// }

	// fmt.Println("Title :", book.Title)
	// fmt.Println("book object : %v", book)

	// mengambil data buku by id
	// err = db.Debug().First(&book, 1).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error saat read data buku")
	// 	fmt.Println("==========================")
	// }

	// fmt.Println("Title :", book.Title)
	// fmt.Println("book object : %v", book)

	// mengambil data semua buku 
	// var books []book.Book

	// err = db.Debug().Find(&books).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error saat read data buku")
	// 	fmt.Println("==========================")
	// }

	// for _, b := range books {
	// 	fmt.Println("Title :", b.Title)
	// 	fmt.Println("book object : %v", b)
	// }
	
	// mengambil data buku dg query string conditions
	// err = db.Debug().Where("title = ?", "buku 1").Find(&books).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error saat read data buku")
	// 	fmt.Println("==========================")
	// }

	// for _, b := range books {
	// 	fmt.Println("Title :", b.Title)
	// 	fmt.Println("book object : %v", b)
	// }
	
	// ==========
	// UPDATE
	// ==========
	// var book book.Book

	// err = db.Debug().Where("id = ?", 2).First(&book).Error
	// if err != nil {
	// 	fmt.Println("============================")
	// 	fmt.Println("Error saat mencari data buku")
	// 	fmt.Println("============================")
	// }
	
	// book.Title = "Revisi Buku 2"
	// err = db.Debug().Save(&book).Error
	// if err != nil {
	// 	fmt.Println("===========================")
	// 	fmt.Println("Error saat update data buku")
	// 	fmt.Println("===========================")
	// }

	// ==========
	// DELETE
	// ==========
	// err = db.Debug().Delete(&book).Error
	// if err != nil {
	// 	fmt.Println("===========================")
	// 	fmt.Println("Error saat delete data buku")
	// 	fmt.Println("===========================")
	// }

	r := gin.Default()

	// grouping by version
	v1 := r.Group("/v1")

	// coba buat routing 
	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)

	r.Run()
	// bisa mengganti default port
	// r.Run(":8888")
}