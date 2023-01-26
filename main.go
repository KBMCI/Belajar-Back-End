package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// coba buat routing 
	r.GET("/", rootHandler)

	r.GET("/hello", helloHandler)

	r.Run()
	// bisa mengganti default port
	// r.Run(":8888")
}

// membuat function handler agar lebih ringkas
func rootHandler(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"title": "ini title",
		"content": "ini content",
	})
}

func helloHandler(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}