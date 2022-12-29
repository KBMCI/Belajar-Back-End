package main;

import "fmt";

func sayHello()  {
	fmt.Println("Hello World");
}

func bilanganKe()  {
	for i := 0; i < 10; i++ {
		fmt.Println("Bilangan ke - ", i);
	}
}

func main() {
	for i := 0; i <= 1; i++ {
		bilanganKe();
	}

	sayHello();
};