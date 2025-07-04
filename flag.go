package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localhost", "database host")
	var username *string = flag.String("username", "root", "Database username")
	var password *string = flag.String("password", "root", "database password")
	var port *int = flag.Int("port", 505, "database port")

	flag.Parse()

	fmt.Println("Username:", *username)
	fmt.Println("Password:", *password)
	fmt.Println("Host:", *host)
	fmt.Println("Port:", *port)

}

// go run flag.go -username=ilham -password=lii -port=100 -host=local
