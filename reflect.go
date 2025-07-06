package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name    string `required:"true" max:"10"`
	Address string `required:"true" max:"20"`
	Email   string `required:"true" max:"30"`
}

type Sample struct {
	Name string `required:"true" max:"10"`
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)

	fmt.Println("Type Name: ", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		var structField reflect.StructField = valueType.Field(i)
		fmt.Println(structField.Name, "with type: ", structField.Type)
		fmt.Println("  required:", structField.Tag.Get("required"))
		fmt.Println("  max:", structField.Tag.Get("max"))
	}
}

func main() {
	readField(Sample{"ilham"})
	readField(Person{"ilham", "bandung", "ilham@gmail.com"})
}
