package main

import "fmt"

func random() interface{} {
	return "ups"
}

func main()  {
	var result interface{} = random()
	var resultString string = result.(string)	
	fmt.Println(resultString)
	
	switch value := result.(type) {
	case string:
		//value auto jadi string
		fmt.Println("string :", value)
	case int:
		//value auto jadi int
		fmt.Println("int :", value)
	default:
		fmt.Println("unknown type")
	}
}