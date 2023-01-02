package main

import "fmt"

func Ups(i int) interface{}  {
	
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func main() {
	
	var test1 interface{} = Ups(1)
	var test2 interface{} = Ups(2)
	var test3 interface{} = Ups(3)
	
	fmt.Println(test1)
	fmt.Println(test2)
	fmt.Println(test3)
}