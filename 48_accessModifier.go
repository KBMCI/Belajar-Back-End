package main

import (
	"Belajar-Back-End/helper"
	"fmt"
)
 // kalo dijava yang biasanya classnya public, private gitu gitu deh pokoknya

func main48() {
	fmt.Println(helper.Application) // tidak error

	// fmt.Println(helper.version)
	// error karena huruf " v " pada kata " version " diawali dengan huruf kecil

	helper.SayHello("Eko")	// tidak error

	// helper.sayGoodbye() 
	// error karena huruf " s " pada kata " sayGoodBye " diawali dengan huruf kecil
}