package main

import (
	"fmt"
	"strings"
)

// PACKAGE STRING
// adalah package yang berisikan function-function untuk memanipulasi tipe data string
// dan dia punya banyak sekali function yang bisa di gunakan, bisa cek dokumentasi golang kalau mau explore

// ada beberapa function yang sering di gunakan
/*
strings.Trim(string, cutset) = memotong cutset di awal dan di akhir
strings.ToLower(string) = membuat semua karakter mejadi lowercase
strings.ToUpper(string) = membuat semua karakter menjadi uppercase
strings.Split(string, separator) = memotong string berdasarkan separator
strings.Contains(string, search) = mengecek apakah string mengandung string lain
strings.ReplaceAll(string, from, to) = mengubah string dari from ke to
dan masih banyak lainnya
*/

// contoh di sini saya langsung pakai pirintln
func main(){
	fmt.Println(strings.Trim("Muhammad Ridho Febrian", "Muhammad")) //maka dia akan menghapus Muhammad dari string
	fmt.Println(strings.Split("Muhammad Ridho febrian", " ")) //dia return nya adalah slice string
	fmt.Println(strings.ToLower("MUHAMMAD RIDHO FEBRIAN")) //maka dia akan mengubah semua string menjadi lowercase
	fmt.Println(strings.ToUpper("muhammad ridho febrian")) //maka dia akan mengubah semua string menjadi uppercase
	fmt.Println(strings.Contains("Muhammad Ridho Febrian", "Ridho")) //maka dia akan mencari string dengan nama Ridho
	fmt.Println(strings.ReplaceAll("Muhammad Brian Febrian", "Brian", "Ridho")) //maka dia akan mengganti Brian menjadi string ridho
}