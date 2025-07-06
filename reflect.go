package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
}

type Sample struct {
	Name string
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)

	fmt.Println("Type Name: ", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		var valueField reflect.StructField = valueType.Field(i)
		fmt.Println(valueField.Name, "with type: ", valueField.Type)
	}

}

func main() {
	readField(Sample{"ilham"})
	readField(Person{"ilham"})
}
