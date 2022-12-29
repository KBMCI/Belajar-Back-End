package main;

import "fmt";

func main() {

	//break

	// if diletakkan sebelum diprint

	for i := 0; i <= 10; i++ {
		if i == 5 {
			break;
		}
		fmt.Println("Perulangan ke - ", i);

	};

	// if diletakkan setelah diprint

	for j := 0; j <= 10; j++ {
		
		fmt.Println("Perulangan ke - ", j);

		if j == 5 {
			break;
		}
	};

	//continue

	for k := 0; k <= 10; k++ {
		if k%2 == 1 {
			continue
		}
		fmt.Println("Perulangan ke - ", k)
	}
};