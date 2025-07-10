package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	value := "Ilham Lii Assidaq"

	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println(encoded)
	// Decode base64 string

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(string(decoded))
	}
	// encoded = base64 menggunakan if else untuk menangani error

}
