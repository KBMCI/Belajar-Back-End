package main

import "fmt"

func main()  {
	for i := 0; i < 5; i++ {
		fmt.Println("angka ke-",i)
	}

	//for mirip while
	count := 1
	for count <= 5 {
		fmt.Println("count ke", count)
		count++
	}

	//for each
	names := [...]string{
		"yanto",
		"agung",
	}
	for _, value := range names { //index bisa dikosongin
		fmt.Println("nama =", value)
	}

	person := map[string]string{
		"name": "yanto",
		"title": "programmar",
	}
	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}