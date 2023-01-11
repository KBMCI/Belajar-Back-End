package main

import (
	"fmt"
	"regexp"
)
func main60() {
	// regexp.MustCompile(string) 			--> Membuat Regexp
	// Regexp.MatchString(string) bool  	--> Mengecek apakah Regexp match dengan string
	// regexp.FindAllString(string, max) 	--> Mencari hasil string yang match dengan maximum jumlah hasil

	var regex *regexp.Regexp = regexp.MustCompile("e([a-z])o")

	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("eto"))
	fmt.Println(regex.MatchString("eDo")) 
	// output false karena huruf D merupakan huruf kapital sedangkan pada regex.MustCompile a-z adalah huruf kecil

	search := regex.FindAllString("eko eka eki ekekekekek eto eyo edo", 10) // n int jika ditulis -1 maka jumlah search tidak terbatas

	fmt.Println(search) 
	// output ygn tercetak hanya huruf depan e kecil dan belakang o kecil  


}