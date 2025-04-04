package main

import (
	"fmt"
	"regexp"
)

// PACKAGE REGEXP
// adalah utilitas golang untuk melakukan pencarian regular expression
// regulra expression di golang menggunakan library dari C, yang dibuat google bersama RE2
// bebrapa function yang sering digunakan
/*
regexp.MustCompile(string) = membuat regexp
regexp.MatchString(string) bool = mengecek apakah regexp cocok dengan string
regexp.FindAllString(string, max) = mencari string yang cocok dengan maximum jumlah hasil
*/

func main(){
	
	regex := regexp.MustCompile(`e([a-z])o`) //a-z itu maksudnya boleh ada karakter atau huruf a-z

	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("elo"))
	fmt.Println(regex.MatchString("eDo"))


	fmt.Println(regex.FindAllString("ridho Ridho RIDHO riho rilo", 10))
}