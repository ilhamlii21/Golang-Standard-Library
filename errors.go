package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("Validation error")
	NotFoundError   = errors.New("Not found error")
)

func GetById(id string) error {

	if id == "" {
		return ValidationError
	}

	if id != "Ilham" {
		return NotFoundError
	}

	return nil

}

func main() {
	err := GetById("1")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("Validation error")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("Not found error")
		} else {
			fmt.Println("Unknown error", err)
		}
	}

}
