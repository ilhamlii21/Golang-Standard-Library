package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Local())

	var utc time.Time = time.Date(2009, time.August, 17, 12, 12, 12, 12, time.UTC)
	fmt.Println(utc)

	// formatter := "2006-01-02 15:04:05"

}
