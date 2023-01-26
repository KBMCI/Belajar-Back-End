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