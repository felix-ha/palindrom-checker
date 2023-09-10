package main

import (
	"fmt"
)

func isPalindrome(s string) bool {
	length := len(s)

	for i := 0; i < length/2; i++ {
		if s[i] != s[length-1-i] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isPalindrome("racecar"))
	fmt.Println(isPalindrome("otto"))

	fmt.Println(isPalindrome("abc"))
}
