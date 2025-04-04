package main

import (
	"fmt"
	"regexp"
)

func main() {
    regex := regexp.MustCompile(`r([a-z])o`)
    fmt.Println(regex.MatchString("ridho")) // Seharusnya true
}