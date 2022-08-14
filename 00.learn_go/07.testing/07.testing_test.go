package main

import "testing"

func add(x, y int) int {
	return x + y
}

func TestAdd(t *testing.T) {
	result := add(2, 3)
	if result != 5 {
		t.Errorf("Expected result to be 5, but got %v", result)
	}
}
