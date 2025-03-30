package main

import (
	"fmt"
	"os"
)

// PACKAGE OS
// golang telah menyediakan banyak package bawaan, salah satunya adalah package os
// package ini berisi fungsionalitas untuk mengakses fitur sistem operasi secara independen(bisa digunakan semua sistem operasi)

func main(){
	args := os.Args //args di gunakan untuk mendapatkan argumen yang dikirim ketika kita menjalankan aplikasi
	//argumen yang dikirim akan menjadi slice di dalam Args ini
	for _, arg := range args{
		fmt.Println(arg)
	}
	//lalu masukkan data nya lewat terminal, kalau mau jadi satu baris gunakan kutip dua ("")

	// contoh lain
	// jika ingin mengambil atau menampilkan nama host atau hostname nya bisa gunakan os.Hostname()
	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("Error", err.Error())
	}
}