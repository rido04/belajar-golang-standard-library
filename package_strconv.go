package main

import (
	"fmt"
	"strconv"
)

//PACKAGE STRCONV(string conversion)
// sebelumnya kita sudah belajar koncersi tipe data, contoh dari int32 ke int64
// strcnov bisa di gunakan untuk konversi tipe data ke yang berbeda, misal dari int ke string atau sebaliknya

// beberapa method yang sering digunakan
/*
strconv.parseBool(string) = mengubah string ke boolean
strconv.parseFloat(string) = mengubah string ke float
strconv.parseInt(string) = mengubah string ke integer
strconv.formatBool(bool) = mengubah boolean ke string
strconv.formatFloat(float, ...) = mengubah float64 ke string
strconv.formatInt(int, ...) = mengubah int64 ke string
*/

func main(){
	// contoh konversi string ke boolean
	boolean, err :=strconv.ParseBool("true") //ini bisa terjadi error kalau string nya bukan nilai boolean itu sendiri(true/false)
	if err ==nil{
		fmt.Println(boolean)
	} else {
		fmt.Println("Error", err.Error())
	}

	// contoh konversi string ke int
	resultInt, err :=strconv.Atoi("9999") //kalau pakai parseInt harus masukkan base int nya, jadi saya pakai alternatifnya yaitu atoi
	if err == nil{
		fmt.Println(resultInt)
	} else {
		fmt.Println("Error", err.Error())
	}

	// contoh kalau mau kebalikannya, yaitu konversi ke string, akan lebih simpel
	boolConv:=strconv.FormatBool(false)
	fmt.Println(boolConv)

	binary := strconv.FormatInt(6969, 2) // disini saya akan ubah ke binary
	fmt.Println(binary)

	intConv := strconv.Itoa(999) // kalau mau konversi ke string
	fmt.Println(intConv)
}