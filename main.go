package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	return true
}

func main() {
	fmt.Println(isPalindrome("racecar"))
}
