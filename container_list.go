package main

import (
	"container/list"
	"fmt"
)

// PACKAGE CONTAINER/LIST
// merupakan implementasi struktur data double linked list di golang
// biasanya di gunakan untuk antrian atau tumpukan

// contoh
func main(){
	// list ini sebenarnya pointer
	data := list.New() //buat dengan menggunakan list.New()
	// lalu gunakan pushback untuk tambahkan ke yang paing belakang
	data.PushBack("Muhammad")
	data.PushBack("Ridho")
	data.PushBack("Febrian")
	// ada banyak method lain selain pushback, tergantung kebutuhan

	// conoth menggunakan pushfront
	data2 := list.New()
	data.PushFront("Febrian")
	data.PushFront("Ridho")
	data.PushFront("Muhammad")

	// gunakan data.Front() untuk mengambil data yang paling depan
	// lalu gunakan next untuk mengambil data selanjutnya
	for e := data.Front(); e != nil; e =e.Next(){
		fmt.Println(e.Value)
	}
	// contoh menggunakan back, dengan previous untuk mengambil data sebelumnya
	for e := data.Back(); e != nil; e =e.Prev(){
		fmt.Println(e.Value)
	}

	for e2 := data2.Front(); e2 != nil; e2 = e2.Next(){
		fmt.Println(e2.Value)
	}
}