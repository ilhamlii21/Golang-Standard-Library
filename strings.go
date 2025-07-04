package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Ilham Lii", "Ilham"))
	fmt.Println(strings.Split("Ilham Lii", " "))
	fmt.Println(strings.ToLower("ILHAM LII"))
	fmt.Println(strings.ToUpper("ilham lii"))
	fmt.Println(strings.Trim("   Ilham Lii   ", " 	"))
	fmt.Println(strings.ReplaceAll("Ilham Lii", "Ilham", "Naruto"))

}
