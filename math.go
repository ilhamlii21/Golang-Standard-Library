package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Ceil(1.35))  // ceil untukm membulatkan ke atas
	fmt.Println(math.Floor(1.60)) // floor untuk membulatkan ke bawah
	fmt.Println(math.Round(1.60)) // round untuk membulatkan ke angka terdekat
	fmt.Println(math.Max(12, 20)) // max untuk mencari nilai maksimum
	fmt.Println(math.Min(12, 20)) // min untuk mencari nilai minimum
}
