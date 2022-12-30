package main

import "fmt"

func main()  {
	for i := 0; i < 10; i++ {
		
		if i % 2 == 0 {
			fmt.Println("genap")
			continue
		}
		fmt.Println("looping ke", i)

		if i == 7 {
			println("berhenti")
			break
		}
	}
}