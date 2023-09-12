package main

import (
	"flag"
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

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func main() {

	s := flag.String("string", "", "String that will be checked if it is a palindrome")
	flag.Parse()

	if isFlagPassed("string") {
		fmt.Println(isPalindrome(*s))
	} else {
		fmt.Println("Argument --string was not found.")
	}

}
