package main

import "fmt"


// menggunakan for loop
func factorialLoop(value int) int  {
	hasil := 1 

	for i := value; i > 0; i-- {
		hasil *= i
	}

	return hasil
}

//menggunakan recoursive function 

func recoursiveFunction(isi int) int {
	if isi == 1 {
		return 1
	}else {
		return isi * recoursiveFunction(isi-1)
	}
}
func main() {
	bilangan := factorialLoop(5)

	fmt.Println(bilangan)

	fmt.Print(5 * 4 * 3 * 2 * 1)

	recoursive := recoursiveFunction(10)

	fmt.Println(recoursive);
}
