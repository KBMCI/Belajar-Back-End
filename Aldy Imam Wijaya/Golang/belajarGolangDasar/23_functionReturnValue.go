package main;

import "fmt";

// yang akan dikembalikan hanya parameter berupa string
func getHello(nama string, nim int64) string  {
	if nama == "" {
		return "Halo Bro"  
	} else {
		return "Halo " + nama
	}
}

func main() {
	hasil := getHello("Aldy", 215150700111039); // nim tidak akan tercetak karena value yang dikembalikan hanya string
	fmt.Println(hasil);

	fmt.Println(getHello("", 0))
};