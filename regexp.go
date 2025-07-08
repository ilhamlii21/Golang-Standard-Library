package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`i([a-z])m`)

	fmt.Println(regex.MatchString("ilham"))
	fmt.Println(regex.MatchString("iam"))
	fmt.Println(regex.MatchString("il"))

	fmt.Println(regex.FindAllString("ilham iam lii", 10))
}
