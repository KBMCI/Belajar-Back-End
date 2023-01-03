package main

import (
	"fmt"
	"strings"
)

func main52()  {
	fmt.Println(strings.Contains("Yanto agung", "Yanto"))
	fmt.Println(strings.Split("yanto agung", " "))
	fmt.Println(strings.Trim("    yanto agung     ", " "))
	fmt.Println(strings.ReplaceAll("yant yanto yanto", "yanto", "agung"))
}