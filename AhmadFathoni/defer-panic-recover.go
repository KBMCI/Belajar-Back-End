package main

import "fmt"

// DEFER
func logging()  {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int)  {
	// defer berarti fungsi yang kita jadwalkan untuk dieksekusi setelah fungsi penampung selesai dieksekusi
	// defer func akan selalu dieksekusi walau terjadi error di function penampung
	// defer usahakan ditaruh di bagian awal function agar bila terjadi error tetap akan dieksekusi 
	defer logging()
	fmt.Println("Run Application")
	result := 10 / value
	fmt.Println("Result", result)

}

// PANIC & RECOVER
func endApp()  {
	// fungsi recover untuk menangkap data panic dan program akan tetap berjalan
	// fungsi recover harus berada setelah panic dieksekusi, untuk case disini ditaruh di func defer 
	// karena func defer akan dieksekusi terakhir meskipun function penampung error
	message := recover()
	if message != nil {
		fmt.Println("Error dg message:", message)
	}
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool)  {
	defer endApp()
	if error {
		// panic brarti sebuah fungsi yang bisa digunakan untuk menghentikan program
		// panic biasanya dipanggil ketika error pada saat program berjalan
		// ketika panic dipanggil, program terhenti, tapi defer function tetap dieksekusi
		panic("APLIKASI ERROR")
	}
	fmt.Println("Aplikasi Berjalan")
}

func main() {
	fmt.Println("== implement defer ==")
	// program akan error jika parameter diisi 0, karena di go angka tdk bisa dibagi 0
	runApplication(10)
	
	fmt.Println("== implement panic ==")
	// ketika tdk error
	// runApp(false)
	// ketika error
	runApp(true)

	fmt.Println("Program masih berjalan lanjut bos")
}