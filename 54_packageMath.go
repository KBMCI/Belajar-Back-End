package main

import (
	"fmt"
	"math"
)
func main54() {
	
	// math.Round(Float64) 	--> Membulatkan float64 keatas atau ke bawah sesuai dengan nilai yang dekat
	// math.Floor(Float64) 	--> Membulatkan float64 kebawah
	// math.Ceil(Float64) 	--> Membulatkan float64 keatas
	// math.Max(Float64) 	--> Mengembalikan nilai float64 paling besar
	// math.Min(Float64) 	--> Mengembalikan nilai float64 paling kecil

	// math.Round
	fmt.Println(math.Round(1.7))
	fmt.Println(math.Round(1.3))

	// math.Floor
	fmt.Println(math.Floor(1.7))

	// math.Ceil
	fmt.Println(math.Ceil(1.3))

	// math.Max & math.Min
	fmt.Println(math.Max(10, 20))
	fmt.Println(math.Min(10, 20))
}