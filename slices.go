package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"ilham", "Lii", "Assidaq"}
	values := []int{1, 2, 3, 4, 5}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(values))

	fmt.Println(slices.Contains(names, "Lii"))

	fmt.Println(slices.Contains(names, "ilham"))
	fmt.Println(slices.Index(names, "ilham"))

}
