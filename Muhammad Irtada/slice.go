package main

import "fmt"

func main() {
	var angka = [...]int{
		1,
		2,
		3,
		4,
		5,
		6,
		7,
		8,
		9,
		10,
		11,
		12,
	}
	var slice1 = angka[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1)) //3
	fmt.Println(cap(slice1)) //8

	// jika melakukan append pada slice yang ujungnya tidak sampai ujung array maka hasil append akan masuk array pada index setelah index ujung slice
	var slice2 = append(slice1, 13)
	fmt.Println(slice1) //[5 6 7]
	fmt.Println(slice2) //[5 6 7 13]
	fmt.Println(angka)  //[1 2 3 4 5 6 7 13 9 10 11 12]

	// jika melakukan append pada slice yang ujungnya merupakan ujung array maka hasil append terdapat pada array baru sehingga tidak mempengaruhi array lama
	slice3 := angka[8:]
	slice4 := append(slice3, 14)
	fmt.Println(slice3) //[9 10 11 12]
	fmt.Println(slice4) //[9 10 11 12 14]
	fmt.Println(angka)  //[1 2 3 4 5 6 7 13 9 10 11 12]

	//perbedaan slice dengan array pada perubahan index. Jika menyamakan 2 array maka hanya terjadi copy data sehingga jika terjadi perubahan data pada array baru tidak mempengaruhi yang lama
	angka2 := angka
	angka2[0] = 0
	fmt.Println(angka) //[1 2 3 4 5 6 7 13 9 10 11 12]
	fmt.Println(angka2) //[0 2 3 4 5 6 7 13 9 10 11 12]

	//jika menggunakan slice maka saat mengubah data pada slice maka juga mengubah data pada array karena by refference
	slice5 := angka2[:]
	slice5[1] = 1
	fmt.Println(slice5) //[0 1 3 4 5 6 7 13 9 10 11 12]
	fmt.Println(angka2) //[0 1 3 4 5 6 7 13 9 10 11 12]
	
	copySlice := make([]int, 2, 2)
	copy(copySlice, slice5)
	fmt.Println(copySlice)

	copySlice2 := append(copySlice, 99)
	fmt.Println(copySlice2)
	fmt.Println(copySlice)
}
