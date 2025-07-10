package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"ilham", "Lii", "Assidaq"})
	_ = writer.Write([]string{"David", "Correnswet", "Superman"})
	_ = writer.Write([]string{"Lex", "Corp", "Luther"})

	writer.Flush()
}
