package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
) // anjir kenapa importnya harus manual ya sksksk :(

func main() {
	
	//go get -u github.com/gin-gonic/gin untuk import 
	route := gin.Default()

	route.GET("/", func(ctx *gin.Context){
		ctx.JSON(http.StatusOK, gin.H{
			"nama": "Aldy Imam WIjaya",
			"deskripsi": "Baik dan ramah",
		})
	})
	
	route.Run()
}