package main

import (
	"fmt"
	"time"
)

func main58()  {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Day())

	utc := time.Date(2022, 10, 10,10,10, 10,10, time.UTC)
	fmt.Println(utc)

	layout := time.RFC3339
	parse, _ := time.Parse(layout, "2022-01-01T15:04:05Z")
	fmt.Println(parse)
}