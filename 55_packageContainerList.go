package main

import (
	"container/list"
	"fmt"
)

func main55() {
	
	ll := list.New()
	ll.PushBack("Imam")
	ll.PushBack("Wijaya")
	ll.PushFront("Aldy") // kata "Aldy" akan muncul sebelum kata "Imam"
	
	fmt.Println(ll.Back().Value)
	fmt.Println(ll.Front().Value)

	// print dari depan
	for i := ll.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	// print dari belakang
	for i := ll.Back(); i != nil; i = i.Prev() {
		fmt.Println(i.Value)
	}
}