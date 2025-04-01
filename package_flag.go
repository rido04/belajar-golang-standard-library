package main

import (
	"flag"
	"fmt"
)

// PACKAGE FLAG
// package flag berisikan fungsionalitas untuk memparsing command line argument
// kalau pakai package os kan biasanya akan dikasih bulat bulat hasilnya
// nah package flag ini kalau kita mau outputnya itu dalam format tertentu

// contoh
func main(){
	// di sini contoh menggunakan string
	// formatnya adalah "nama flag", "nilai default nya", "lalu deskripsinya"
	host := flag.String("host", "localhost", "Host Database anda")
	username:= flag.String("username", "username", "Username database anda")
	password := flag.String("password", "password", "Password database anda")
	port := flag.Int("port", 6969, "Port database anda")
	is_db:= flag.Bool("asli", true, "contoh kalau pakai boolean")
	// perlu di perhatikan, tipe data yang dikembalkan adalah pointer
	// jadi disini data yang dikembalikan adalah pointer string

	// lalu gunakan flag.Parse()
	flag.Parse()

	// karena data yang dikembalikan adalalh pointer, jadi harus menggunakan operator asterisk(*) di depannya
	fmt.Println("Host",*host,"Username", *username,"Password", *password,"Port", *port,"asli ga", *is_db)

	// kalau tidak menggunakan asterisk maka data yang di print adalah si pointer nya
	fmt.Println(host, username, password)

	// ketika menjalankan aplikasi, golang akan menjalankan default value yang ada di flag
	// kalau mau kirim data ke value nya cukup menambahkan mines(-) ketika menjalankan aplikasi

	// contoh
	// go run package_flag.go -host=127.0.0.1 -username=myuser -password=mypass -port=5432 -asli=false
	// NOTE :
	// pastikan ketika mengirim data ke value harus sesuai dengan hurud besar dan kecilnya si flag
}