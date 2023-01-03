package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main56()  {
	var data *ring.Ring = ring.New(5)

	for i := 0; i < data.Len(); i++ {
		data.Value = "data " + strconv.Itoa(i)
		data = data.Next()
	}

	data.Do(func(a any) {
		fmt.Println(a)
	})
}