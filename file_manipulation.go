package main

import (
	"bufio"
	// "fmt"
	"io"
	"os"
)

// FILE MANIPULATION
// di package os, terdapat file management
// saat kita mmebuat atau membaca file menggunakan package os, struct file meruapakan implementasi dar io,Reader dan io.Writer
// makanay kita bisa melakukan baca dan tulis terhada file tersebut menggunakan package io/bufio

// OPEN FILE
// unutk membuat atau membaca file, kita bisa menggunakan os.OpenFile(name, flag, permission)
// name berisikan nama file, bisa absolut atau relatif/local
// flag merupakan penanda file, pakah untuk membaca, menulis, dll
// permission merupakan izin atau hak akses yang diperlukan ketika membuatf file
// kalau di linux biasa nya rwrr, rdwr dsb

// contoh kalo buat file
func buatFile(nama string, pesan string) error{
	// gunakan os.OpenFile untuk membuat file baru
	// kalau lebih dari satu bisa tambahkan pipe (|) diantara flag
	file, err := os.OpenFile(nama, os.O_CREATE|os.O_WRONLY, 0666) //lalu gunakan os.O_CREATE untuk membuat file baru, os.O_WRONLY untuk hanya menulis ke file, dan 0666 untuk permission
	if err != nil {
		return err
	}
	defer file.Close() //gunakan file.Close() untuk menutup file setelah selesai
	file.WriteString(pesan) //lalu gunakan file.WriteString untuk menulis ke file
	return nil
}

// contoh kalau baca file
func bacaFile(nama string)(string, error){
	// pertama kita buat untuk buka file nya
	file, err := os.OpenFile(nama, os.O_RDONLY, 0666) //gunakan os.O_RDONLY untuk hanya membaca file, pastikan filenya sudah ada
	if err != nil {
		return "", err
	}
	defer file.Close()

	// buat reader untuk membacar filenya
	reader := bufio.NewReader(file) //inputnya dari file yang sudah kita buka
	var message string
	for {
		line, _, err := reader.ReadLine()
		message += string(line) + "\n"
		if err == io.EOF{
			break
		}
	}
	return message, nil
}

	// contoh membaca dan menambah file
	func tambahFile(nama string, pesan string) error{
		file, err := os.OpenFile(nama, os.O_RDWR|os.O_APPEND, 0666) //gunakan os.O_RDWR untuk membaca dan menulis file, os.O_APPEND untuk menambah isi file
		if err != nil{
			return err
		} 
		defer file.Close()
		file.WriteString(pesan) //jadi dia tidak akan menimpa isi file yang sudah ada, tapi menambahkannya di akhir
		return nil
	}


func main(){
	// contoh penggunaannya, bisa langsung panggil si functionnya
	// buatFile("error.log", "ini adalaha Log error")
	// result,_ := bacaFile("error.log")
	// fmt.Println(result)
	tambahFile("error.log", "\nini adalah log error tambahan2")
}