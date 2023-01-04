package main

import "fmt";

func main() {
	type NIM int64;
	type status bool;

	var nim NIM = 215150700111021;
	fmt.Println(nim);

	var status2 status = true;
	fmt.Println(status2);
};