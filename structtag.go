package main

import (
	"fmt"
	"reflect"
)

// ADAPun STRUCTAG
// yaitu tag di dalam field
// dan ini banyak sekali di gunakan untuk membuat json, validation

// contoh
type Anjay struct {
	// struktur tag adalah nama tag:value tag
	Name string `required:"true" max:"10"` //gunakan backtik atau backquote untuk membuat tag
}

type Orang struct{
	Name string `required:"true" max:"10"` //gunakan backtik atau backquote untuk membuat tag
	Address string `required:"true" max:"10"` //gunakan backtik atau backquote untuk membuat tag
	Email string `required:"true" max:"10"` //gunakan backtik atau backquote untuk membuat tag
}

// contoh implementasi untuk validation
func isValid(value interface{}) (result bool){
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" { //kalau tag "required" nya true
			// maka wajib ada datanya
			
			data :=  reflect.ValueOf(value).Field(i).Interface()
			result = data != ""
			if result == false{
				return	result //disini saya membuat agar semua data harus di isi, jadi tidak ada data yang kosong
			}
		}
	}
	return result
}
func main(){
	anjay := Anjay{"Mabar"}
	anjayType := reflect.TypeOf(anjay)
	structfield := anjayType.Field(0)
	required := structfield.Tag.Get("required")//gunakan Tag.Get() untuk mengambil tagnya
	required2 := structfield.Tag.Get("max")//gunakan Tag.Get() untuk mengambil tagnya

	fmt.Println(required, required2)

	orang := Orang{
		Name: "Muhammad Ridho febrian",
		Address: "Pakulonan Barat",
		// Email: "Konoha",
	}

	fmt.Println(isValid(orang)) //gunakan function validasi nya

}