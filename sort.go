package main

import (
	"fmt"
	"sort"
)

//PACKAGE SORT
// adalah package yang berisikan utilitas untuk proses pengurutan
// agar data bisa di urutkan, kita harus mengimplementasikan kontrak di interface menggunakan sort.Interface

type Interface interface{
	Len() int //harus punya len, yaitu panjang dari elemennya
	Less(i, j int) bool //untuk mengecek kurang dari atau tidak
	Swap(i, j int) //untuk menukar posisi
}

// buat struct user
type User struct{
	Name string
	Age int
}

// lalu buat type declaration untuk user sesuai kontrak agar bisa di urutkan dengan mudah
type UserSlice []User //ini adalah slice untuk struct User

// kalau mau pakai sort, harus buat kontrak dlu dengan slice nya
func (userSlice UserSlice) Len() int{ //mwnggunakan Len() untuk menentukan jumlah panjang nya
	//tidak butuh operator arterisk karena defaultnya slice itu sudah pointer
	return len(userSlice)
}

// untuk lessnya
func(userSlice UserSlice) Less(i, j int) bool{
	return userSlice[i].Age < userSlice[j].Age
}

// lalu swapnya
func (userSlice UserSlice) Swap(i, j int){
	userSlice[i], userSlice[j] = userSlice[j], userSlice[i]
}

func main(){
	user := []User{
		{"Muhammad", 30},
		{"Ridho", 35},
		{"Febrian", 35},
		{"Ulala", 20},
	}

	sort.Sort(UserSlice(user)) //konversi slice [User] ke UserSlice nya

	fmt.Println(user)
}