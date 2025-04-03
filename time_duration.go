package main

import (
	"fmt"
	"time"
)

// adapun DURATION
// biasanya di butuhkan ketika kita butuh data berupa durasi
// package time punya type duration, yang sebenernya adalah alias untuk int64
// banyak method yang bisa kita gunakan untuk memanipulasi duration

// contoh

func main(){
	duration1 := time.Second * 100 //100 adalah detiknya
	duration2 := time.Millisecond * 10 //10 adalah milidetiknya
	duration3 := duration1 - duration2

	fmt.Println("duration1: %d\n", duration3)
	}