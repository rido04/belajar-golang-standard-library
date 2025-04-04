package main

import (
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

// PACKAGE ENCODING
// golang juga punya package untuk encode dan decode
// contoh algoritma yang populer adalah base64, csv, dan json

func main(){
	// conoth kode base64
	value := "Muhammad Ridho Febrian"

	// gunakan EncodeToString unutuk encode
	encoded := base64.StdEncoding.EncodeToString([]byte(value)) //ini akan dikonversi menjadi byte slice
	fmt.Println(encoded)

	// gunakan DeocedeString untuk decode
	decoded, err := base64.StdEncoding.DecodeString(encoded) //bisa kirim error juga, karena bisa saja data nya bukan base64
	if err != nil{
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(string(decoded))
	}

	// contoh kode csv reader
	csvString := "Muhammad,Ridho,Febrian\n" + //pakai koma untuk pemisah kolom
		"Febrian,Muhammad,Ridho\n" + // \n untuk enter baris baru
		"Ridho,Muhammad,Febrian\n"

	// buat reader nya dengan csv.NewReader, 
	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read() //gunakan Read() untuk membaca csv
		if err == io.EOF { //EOF menandakan akhir dari file (end of file)
			break
		}
		fmt.Println(record) //akan mengembalikan slice string
	}

	//  contoh kode csv writer
	writer := csv.NewWriter(os.Stdout) //disini saya write ke terminal dulu
	// satu write berarti satu baris
	// gunakan underscore (_) untuk menghiraukan error nya untuk saat ini
	_ = writer.Write([]string {"Muhammad", "Ridho", "Febrian"}) //ini akan menulis ke stdout (ke layar)
	_ = writer.Write([]string {"Febrian", "Muhammad", "Ridho"}) //ini juga akan menulis ke stdout (ke layar)
	_ = writer.Write([]string {"Ridho", "Muhammad", "Febrian"}) //ini juga akan menulis ke stdout (ke layar)

	writer.Flush() //flush untuk menulis semua data ke stdout (ke layar) atau mengirim semua perubahannya
	// hasilnya adalah csv
}