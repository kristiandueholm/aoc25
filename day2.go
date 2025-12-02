package main

import (
	"fmt"
	"regexp"
)

func main() {
	teststring := "1212"
	r := regexp.MustCompile(`([1-9]\d*)\1`)
	match := r.MatchString(teststring)
	fmt.Printf("Mathces: %t", match)
}
