package main

import (
	"fmt"
	"sort"
)

type User struct {
	Nama string
	Umur int
}

type UserSlice []User

func (userSlice UserSlice) Len() int           { return len(userSlice) }
func (userSlice UserSlice) Less(i, j int) bool { return userSlice[i].Umur < userSlice[j].Umur }
func (userSlice UserSlice) Swap(i, j int)      { userSlice[i], userSlice[j] = userSlice[j], userSlice[i] }

func main57() {

	 user := []User{
		{"Aldy", 20},
		{"Faza", 19},
		{"Toni", 21},
		{"Bariq", 22},

	 }

	 sort.Sort(UserSlice(user))

	 fmt.Println(user)
}