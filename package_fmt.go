package main

import "fmt"

// PACKAGE FORMAT
// sebelumnya kita sudah sering menggunakan fmt atau format dengan menggunakan println
// selain println, masih banyka function yang terdapat di package fmt, contohnya banyak digunakan untuk melakukan format

func main(){
	// contoh menggunakna println yang biasa kita gunakan
	fmt.Println("Hello World")
	// jika menggunakan println, besar kemungkinan akan memanggil atau menggunakan banyak line, sehingga tidak efisien
	fmt.Println("Muhammad Ridho febrian")

	firstName := "Muhammad"
	middleName := "Ridho"
	lastName := "Febrian"

	// conoth menggunakan println juga
	fmt.Println("Hello '", firstName, middleName, lastName, "'")

	// contoh menggunakan printf
	// gunakan %s sebanyak 3 kali karena akan memanggil 3 variabel di atas
	fmt.Printf("Hello '%s %s %s'\n", firstName, middleName, lastName) // !\n adalah enter, karena printf tidak memakai enter seperti println
	// setiap variabel di atas (firsName, middleName, lastName) akan di substitusi atau di kirim ke dalam %s, jadi sesuaikan jumlah %s nya sesuai jumlah variabel yang ada atau variabel yang dipanggil
}