package main

import (
	"fmt"
	"time"
)

// PACKAGE TIME
// adalah package yang berisikan fungsi untuk management waktu di bahasa golang
// bebrapa function yang sering digunakan di package time
/*
time.Now() = untuk mendapatkan waktu sekarang
time.Date(...) = untuk membuat waktu
time.Parse(layout, string) = untuk memparsing waktu dari string
*/

// contoh
func main(){
	now := time.Now()
	fmt.Println(now.Local()) //gunakan .Local() untuk mendapatkan waktu sesuai region atau negara

	utc := time.Date(2012, time.February, 10, 23, 0, 0, 0, time.UTC) //format nya tahun, bulan, hari, jam, menit, detik, nanodetik, location
	fmt.Println(utc.Local())

	// formatting nya harus mengikuti aturan dar golang time nya
	formatter := "2006-01-02 15:04:05" // Format waktu Go
    value := "2020-10-10 10:10:10"
	parse, err := time.Parse(formatter, value) //RFC3339 adalah format baku, bisa di cek di dkoumentasi atau ketik RFC3339 nya
	if err != nil{
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(value)
	}
	// bisa juga ambil sesuai yang diinginkan
	fmt.Println(utc.Year())
	fmt.Println(now.Second())
	fmt.Println(parse.Hour())
}