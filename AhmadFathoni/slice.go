package main

import "fmt"

func main() {
	// " ... " berarti panjang array tidak terdefinisikan yang menjadikan panjang array bisa fleksibel
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	// cara membuat slice dengan reference array yang sudah ada
	var slice1 = months[4:7]
	fmt.Println(slice1)
	// len : panjang dari slice
	fmt.Println(len(slice1))
	// cap : kapasitas dari slice
	fmt.Println(cap(slice1))

	// saat merubah array yang mereference ke slice yang dibuat, maka slice akan berubah juga
	// months[5] = "index 5 Dirubah"
	// fmt.Println(slice1)
	
	// saat merubah slice yang mereference dari array, maka array akan berubah juga
	// slice1[0] = "Index 0 Slice Dirubah"
	// fmt.Println(months)

	var slice2 = months[10:]
	fmt.Println(slice2)

	// fungsi append jika kapasitas sudah melebihi batasnya, maka akan membuat array baru dan jika kita merubah array dari slice3 yang
	// append dari slice2, maka tidak akan merubah array atau slice yang direference, contoh bisa di kode nomor 48 dan 49
	// dan jika kapasitas tidak melebihi batas, maka perubahannya akan mereplace data array dan slice yang di reference
	var slice3 = append(slice2, "Toni")
	fmt.Println(slice3)
	slice3[1] = "Ubah Desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	// fungsi make untuk membuat slice baru dengan parameter([]tipeData, length, capacity)
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Ahmad"
	newSlice[1] = "Fathoni"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	// fungsi copy untuk mengcopy data slice dengan parameter(wadah destinasi, slice mana yang ingin di copy)
	// pastikan parameter len(panjang) harus sama dengan slice yang ingin dicopy, 
	// jika beda, data yang ditampilkan akan sesuai dengan len(panjang) yang ditentukan
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	// perbedaan array dan slice
	iniArray := [...]int{1,2,3,4,5}
	iniSlice := []int{1,2,3}
	// beda hanya di deklarasi panjang index

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}