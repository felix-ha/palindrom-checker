package main

import "testing"

func TestExampleTrueEven(t *testing.T) {
	result := isPalindrome("anna")
	if result == false {
		t.Errorf("Expected result to be true, got false")
	}
}

func TestExampleTrueOdd(t *testing.T) {
	result := isPalindrome("rentner")
	if result == false {
		t.Errorf("Expected result to be true, got false")
	}
}

func TestExampleFalseEven(t *testing.T) {
	result := isPalindrome("abcd")
	if result == true {
		t.Errorf("Expected result to be false, got true")
	}
}

func TestExampleFalseOdd(t *testing.T) {
	result := isPalindrome("abcde")
	if result == true {
		t.Errorf("Expected result to be false, got true")
	}
}
