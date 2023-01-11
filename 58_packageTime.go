package main

import (
	"time"
	"fmt"
)
func main58() {
	// time.Now() 					--> Untuk mendapatkan waktu saat ini
	// time.Date(...) 				--> Untuk membuat waktu
	// time.Parse(layout, string) 	--> Untuk memparsing waktu dari string

	// time.Now
	now := time.Now()

	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Date())  
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	// time.Date
	utc := time.Date(2023, 1, 11, 21, 25, 59, 0, time.UTC)

	fmt.Println(utc)

	// time.Parse
	layout := time.RFC3339
	parse, _ := time.Parse(layout, "2023-01-11")

	fmt.Println(parse)
}