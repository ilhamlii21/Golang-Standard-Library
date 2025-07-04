package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	var data *ring.Ring = ring.New(5)

	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.Itoa(i)
		data = data.Next()

	}

	// data.Value = "Ilham 1"

	// data = data.Next()
	// data.Value = "Ilham 2"

	// data = data.Next()
	// data.Value = "Ilham 3"

	// data = data.Next()
	// data.Value = "Ilham 4"

	// data = data.Next()
	// data.Value = "Ilham 5"

	data.Do(func(value any) {
		fmt.Println(value)

	})

}
