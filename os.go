package main

import (
	"fmt"
	"os"
)

func main50() {
	// args berupa slice array dimana index 0 berisi alamat file dijalankan
	// apabila dijalankan dengan go run os.go yanto -> maka yanto akan menjadi index 1
	args := os.Args
	fmt.Println(args)

	name, err := os.Hostname()
	if err != nil {
		fmt.Println("error", err.Error())
	} else {
		fmt.Println("hostname :", name)
	}
}
