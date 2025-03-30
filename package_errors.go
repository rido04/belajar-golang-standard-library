package main

import (
	"errors"
	"fmt"
)

// PACKAGE ERRORS
// biasanya kita membuat error dengan function errors.New()
// sebenarnya masih banyak lagi yang bisa kita lakukakn dengan package inin=, contohnya ketika kita ingin membuat bebrapa value error yang berbeda

// biasanya error digolang dibuat seperti ini
var ( //di simpan dalams satu variabel atau constant
	ValidationError = errors.New("Validasi Error")
	NotFoundError = errors.New("Data tidak ditemukan")
	// jadi errornya di buat terlebih dahulu
)

func GetById(id string) error{
	if id == ""{
		return ValidationError
	} 
	
	if id != "Ridho"{
		return NotFoundError
	}

	return nil
}

// ada juga package dari golang untuk mengecek error, kita bisa gunakan error.Is() untuk mengecek jenis type errornya
// biasanya ini digunakan kalau error yang di buat bukan custom error dan memiliki type error yang sama
func main(){
	err := GetById("")
	if err != nil{
		if errors.Is(err, ValidationError){
			fmt.Println("Validasi Error")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("Data tidak ditemukan")
		} else {
			fmt.Println("unknown error")
	}
}
}