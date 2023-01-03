package main

import (
	"fmt"
	"strconv"
)

func main53()  {
	boolean, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println(boolean)
	}

	number, err := strconv.ParseInt("1000000", 10, 64) // par 1 angkanya, par 2 jenisnya (biner, desimal, heksadesimal), par 3 jenis int nya (int32, int64)
	if err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println(number)
	}

	value :=  strconv.FormatInt(1000000, 10)
	fmt.Println(value)

	// itoa atau atoi bisa langsung
	valueString:= strconv.Itoa(1000000)
	fmt.Println(valueString)
}