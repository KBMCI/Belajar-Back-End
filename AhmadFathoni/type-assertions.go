package main

import "fmt"

func random(angka int) interface{} {
	if angka == 1 {
		return "ups"
	}else if angka == 2 {
		return 2
	}else {
		return true
	}
}

func main()  {
	var result interface{} = random(3)
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