package main
import "fmt"

func endApp()  {
	fmt.Println("Aplikasi sudah berjalan")
}

func startApp(status bool)  {
	defer endApp()
	if status {
		panic("Terdapat error pada aplikasi")
	}
	fmt.Println("Yeay aplikasi berjalan lancar")
}
func main() {
	startApp(false) // tidak error
	startApp(true) // error
}