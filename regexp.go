package main

import (
	"fmt"
	"regexp"
)

func main()  {
	var regex *regexp.Regexp = regexp.MustCompile("e([a-z])o")

	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("edo"))
	fmt.Println(regex.MatchString("eTo"))

	search := regex.FindAllString("eko eka edo eto eyo eki", -1) // -1 mencari sebanyak2nya
	fmt.Println(search)
}