package main

import (
	"fmt"
	"slices"
)

// PACKAGE SLICE
// di golang versi terbaru, ada fitur bernama generic programming
// yaitu fitur dimana kita bisa membuat parameter dengan tipe yang bisa berubah tanpa harus menggunakan interface kosong
// package slices adalah asalah satu package yang menggunakan fitur ini
// package slice digunakan untuk memanipulasi data slice

func main(){
	names := []string{"Muhammad", "Ridho", "Febrian"}
	values := []int{100, 95, 90, 85, 80}

	// gunakan slices. untuk mrnggunakan package slices
	fmt.Println(slices.Min(values)) //akan mencari data terkceil dari slice int di atas
	fmt.Println(slices.Max(values)) //akan mencari data terbesar dari slice int di atas
	fmt.Println(slices.Contains(names, "Muhammad")) //akan mencari apakah ada data Muhammad di slice names
	fmt.Println(slices.Contains(names, "Joko Nawa")) //akan mencari apakah ada data Joko Nawa di slice names
	fmt.Println(slices.Index(names, "Ridho")) //akan mencari index dari data Ridho di slice names
	fmt.Println(slices.Grow(names, 5)) //akan menambah kapasitas dari slice names menjadi 5
	fmt.Println(slices.ContainsFunc(names, func(s string) bool {
		return s == "Muhammad"
	}))

}