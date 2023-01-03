package main

import (
	"Belajar-Back-End/helper"
	"fmt"
)

func main48()  {
	helper.SayHello("yanto")
	// helper.sayGoodbye("yanto") error karena private
	fmt.Println(helper.Application)
}