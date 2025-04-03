package main

import (
	"fmt"
	"reflect"
)

// PACKAGE REFLECT
// reflection digunakan untuk melihat struktur kode kita pada saat aplikasi sedang berjalan
// untuk membaca sebuah struktur metadata dari suatu object di golang
// di golang kita bias gunakan package reflect
// sangat berguna untuk membuat library general sehingg mudah digunakan

// contoh
type Sample struct{
	Name string
	Age int
}
type Person struct{
	Name string
}

func readField(value any){
	valueType := reflect.TypeOf(value) //gunakan TypeOf()
	fmt.Println("type name", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		valueField := valueType.Field(i)
		fmt.Println(valueField.Name, "with type", valueField.Type)
	}
}
func main(){
	// Sample := Sample{"Ridho", 21}
	// sampleType := reflect.TypeOf(Sample) // disini saya gunakan TypeOf()
	// structField := sampleType.Field(1) //lalu gunakan .Field(index ke 0) untuk mengambil nama fieldnya, sesuaikan dengan index yang mau di ambil

	// fmt.Println(structField.Name)
	readField(Sample{"Ridho", 21})
	readField(Person{"ridho"})
}