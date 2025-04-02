package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

// PACKAGE CONTAINER RING
// merupakan implementasi struktur data cricular list
// circular list adalah struktur data ring, dimana diakhir element kana mengembalikan element pertama atau elmeent awal
// hati hati jika ingin menggunakan loop, karena bisa terjadi infinite loop, jadi hati hati jika menggunakan next

// contoh
func main(){
	// gunakan ring.New()
	data := ring.New(5) // lalu tentukan parameternya mau buat berapa di dalam circular list nya
	for i := 0; i < data.Len(); i++{
		data.Value = "Value ke " +strconv.FormatInt(int64(i), 10) //karena tidak bisa menggunakan i, jadi harus di konversi dulu menggunakan strconv
		data = data.Next()
	}

	// jika ingin melakukan sesuatu terhadap datanya bisa gunakan Do(func(any))
	data.Do(func (value interface{})  {
		fmt.Println(value)
	})
}