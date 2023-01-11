package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Tambah framework gin dengan mengetikkan
	// go get -u github.com/gin-gonic/gin

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
	router.GET("/", rootHandler)
	router.GET("/books/:id", booksHandler) // :id -> sebuah variable, untuk menandakan variable menggunakan ":"
	router.GET("/query", queryHandler)

	// Method POST
	router.POST("/books", postBooksHandler)

	// menjalankan server
	router.Run() // -> port default 8080
	// router.Run(":8888") -> parameter untuk mendefenisikan port yang digunakan

}

// Membuat handler
// nama fungsi disesuaikan dengan nama route jika /hello nama : helloHandler
func rootHandler(c *gin.Context) {
	// c.JSON digunakan untuk mengirim data
	c.JSON(http.StatusOK, gin.H{
		"name": "Wisnu",
		"bio":  "Aktor film meteor",
	})
}

func booksHandler(c *gin.Context) {
	//untuk menangkap variable menggunakan param
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func queryHandler(c *gin.Context) {
	//untuk menangkap query string menggunakan Query()
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"price": price,
	})
}

// untuk menerima data yang dirikim melalui post dapat diterima menggunakan struct
type BookInput struct{
	// Pemberian nama property harus sama dengan data yang di POST

	// json: -> Untuk mengidentifikasi nama pada json API
	// binding:"required" -> untuk menandakan field tersebut harus diisi
	// Jika saat POST data yang terdapat required pada attributenya tidak memberikan nilai maka akan terjadi error
	Title string `json:"title" binding:"required"`
	Price int `json:"price" binding:"required"`
	SubTitle string `json:"sub_title"` // jika beda dapat menggunakan cara ini
	// SubTitle string -> tidak bisa digunakan
	// Sub_Title string -> bisa digunakan
}

func postBooksHandler(c *gin.Context) {
	var bookInput BookInput

	//Memasukkan data post ke bookInput
	err := c.ShouldBindJSON(&bookInput) // menggunakan pointer karena kalo tidak maka bookInput yang menerima data hasil post hanya secara local variable
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bookInput)

	c.JSON(http.StatusOK, gin.H{
		"title": bookInput.Title,
		"price": bookInput.Price,
		"sub_title": bookInput.SubTitle, 
	})
}
