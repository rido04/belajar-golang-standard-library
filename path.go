package main

import (
	"fmt"
	"path"
	"path/filepath"
)

// PACKAGE PATH
// digunakan untuk memanipulasi data path seperti path di url atau pasth file system
// secara default package path menggunakan slash sebagai karakter pathnya, oleh karena itu cocok dengan data url
// namun jika ingin memanipulasi data path file system khususnya windows karena windows menggunakan backslash, perlu menggunakan package path/filepath

// contoh path
func main(){
	// perlu diingat, ini adalah bikin path, bukan bilin folder
	fmt.Println(path.Dir("hello/world.go")) //ini adalah ambil direktori (yang hello)
	fmt.Println(path.Base("hello/world.go")) // ini adlaah ambil base (yang world.go)
	fmt.Println(path.Ext("hello/world.go")) //ini adalah ambil extension (yang .go)
	fmt.Println(path.Join("hello", "world", "main.go"))//ini adlaah join beberapa path (yang hello/world/main.go)

	// CONTOH FILE PATH
	// bedanya dia bisa tau apa sistem operasi yang digunakan
	fmt.Println(filepath.Dir("hello/world.go")) //ini adalah ambil direktori (yang hello)
	fmt.Println(filepath.Base("hello/world.go")) // ini adlaah ambil base (yang world.go)
	fmt.Println(filepath.Ext("hello/world.go")) //ini adalah ambil extension (yang .go)
	fmt.Println(filepath.Join("hello", "world", "main.go"))//ini adlaah join beberapa path (yang hello/world/main.go)
	fmt.Println(filepath.IsAbs("hello/world.go")) //ini adalah untuk mengecek apakah path absolut atau tidak (false karena di windows biasanya dari depan dimulai dari nama drive)
	fmt.Println(filepath.IsLocal("hello/world.go")) //ini adalah untuk mengecek apakah path lokal atau tidak (true)

}