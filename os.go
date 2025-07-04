package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	for _, arg := range args {
		fmt.Println(arg)
	} //* C:\Users\ASUS\AppData\Local\Temp\go-build1603021576\b001\exe\os.exe adalah sebuah package yang digunakan untuk mengakses sistem operasi

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("Error", err.Error)
	}
}
