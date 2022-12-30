package main

import "fmt"

func main()  {
	var names [3]string
	names[0] = "yanto"
	names[1] = "agung"
	names[2] = "surya"

	fmt.Println(names)

	var umur = [3]int{
		20,
		30,
		40,
	}
	fmt.Println(umur)
	fmt.Println(len(umur))
}