package main

import "fmt";

func katakanHaiKepada(namaDepan string, namaBelakang string, nim int64) {
	fmt.Println("Hai", namaDepan, namaBelakang, "dengan nim ", nim)
}

func main() {
	katakanHaiKepada("Aldy", "Imam", 215150700111039)
};