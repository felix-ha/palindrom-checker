package main

import "testing"

func TestExample(t *testing.T) {
	result := isPalindrome("anna")
	if result == false {
		t.Errorf("Expected result to be true, got false")
	}
}
