package main

import "fmt"

type Blacklist func (nama string) bool // menggunakan type declaration (ada di bab nomer 10)

func registrasiAkun(nama string, daftarBlacklist Blacklist)  {
	
	if daftarBlacklist(nama) {
		fmt.Println("Sayang sekali ternyata", nama, " telah diblacklist")
	}else {
		fmt.Println("Halo", nama)
	}
}

func main() {

	// opsi 1 
	blacklist := func(nama string) bool {
		return nama == "Admin"
	}

	registrasiAkun("Admin", blacklist)
	registrasiAkun("Aldy", blacklist)

	//opsi 2

	registrasiAkun("Imam", func(nama string) bool {
		return nama == "Imam"
	 })
	
	 registrasiAkun("Wijaya", func(nama string) bool {
		return nama == "Imam"
	 })


	
}