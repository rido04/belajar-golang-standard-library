package main

import (
	"fmt"
	"math"
)

// PACKAGE MATH
// merupakan package yang berisi opearsi matematika
// bebrapa function yang sering digunakan
/*
math.Round(float64) = membulatkan float 64 keatas atau kebawah, sesuai dengan yang paling dekat
math.Floor(float64) = membulatkan float 64 ke bawah
math.Ceil(float64) = membulatkan float64 ke atas
math.Max(float64, float64) = mengembalikan nilai float paling besat
math.Min(float64, float64) = mengembalikan nilai float paling kecil
*/

// contoh
func main(){
	fmt.Println(math.Round(6.7))
	fmt.Println(math.Floor(6.6))
	fmt.Println(math.Ceil(6.3))
	fmt.Println(math.Max(6.7, 6.3))
	fmt.Println(math.Min(6.7, 6.3))
}