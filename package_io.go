package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// PACKAGE IO
// package io adalah singkatan dari input output
// merupakan fitur golang yang di gunakan sebagia standar input output
// semua mekanisme input output di golang pasti menggunakan package ini
// isinya sebenernya adalah interface io

// untuk membaca input golang punya interface namanya Reader di package io(kontraknya)
// ada juga namanya writer untuk menulis output, juga menggunakan kontrak interface pada package io

// bagaimana implementasi io nya?
// penggunaan dari io sendiri di goolang terdapat banyak package, contoh sebelumnya adalah csv dan reader
// karena package io sebenernya hanya kontrak untuk io, untuk implementasinya kita harus lakukan sendiri
// untungnya golang menyediakan package untuk mempermuda implementasi io nya, yang menggunakan package bufio namanya

// PACKAGE BUFIO
// adalah singkatan dari buffered io
// adalah package yang di gunakan untuk membuat data io seperti reader dan writer

// contoh kode reader
func main(){
	// gunakan strings.NewReader() untuk membuat reader baru, saya gunakan string di sini
	input := strings.NewReader("ini adalah string yang panjang \nDibuat oleh Muhammmad Ridho Febrian\n")
	reader := bufio.NewReader(input) //lalu kita gunakan bufio.NewReader() untuk membuat reader baru baru input strings di atas

	for {
		// di bufio ada funcgsio ReadLine() untuk membaca satu baris
		line, _, err := reader.ReadLine() //gunakan ReadLine() untuk membaca satu baris
		if err == io.EOF{
			break
		}
		fmt.Println(string(line))
	}


	// contoh kode writer
	writer := bufio.NewWriter(os.Stdout) // gunakan bufio.NewWriter() untuk membuat writer baru, saya gunakan os.Stdout di sini untuk melihat di terminal
	writer.WriteString("ini adalah string yang panjang \nDibuat oleh Muhammmad Ridho Febrian\n") // disini saya gunakan writeString() untuk menulis string ke writer
	writer.WriteString("Anjay keren belajar golang") //writeString() ini juga bisa error, jadi bisa gunakan if atau for, atau gunakan ignore
	writer.Flush() // flush untuk menulis ke output, jika tidak di flush tidak akan muncul di output
}