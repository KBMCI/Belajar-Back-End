package main

import (
	"container/list"
	"fmt"
)

func main55()  {
	data := list.New()

	data.PushBack("yanto")
	data.PushBack("agung")
	data.PushFront("budi")

	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)
	fmt.Println(data.Front().Next().Value)

	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}
}