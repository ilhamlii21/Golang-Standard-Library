package main

import (
	"container/list"
	"fmt"
)

func main() { // penerapan double linked list
	data := list.New()
	data.PushBack("Ilham")
	data.PushBack("Lii")
	data.PushBack("Assidaq")
	// Remove the first element

	var head *list.Element = data.Front()
	fmt.Println("Head:", head.Value)

	next := head.Next()
	fmt.Println(next.Value)

	next = next.Next()
	fmt.Println(next.Value)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	} //for e untuk perulangan setiap elemen
}
