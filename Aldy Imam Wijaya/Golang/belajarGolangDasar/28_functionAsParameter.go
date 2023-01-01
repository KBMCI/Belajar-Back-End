package main
import "fmt"

type Filter func(string) string // menggunakan type declaration (ada di bab nomer 10)

func sayHelloWithFilter(nama string, filter Filter) {
	namaFilter := filter(nama)
	fmt.Println("Halo ",namaFilter)
}  

func spamFilter(nama string) string {
	if nama == "Anjing" {
		return "***"
	}else {
		return nama
	}
}

func main() {
	sayHelloWithFilter("Eko", spamFilter)

	filter := spamFilter

	sayHelloWithFilter("Anjing", filter)
}