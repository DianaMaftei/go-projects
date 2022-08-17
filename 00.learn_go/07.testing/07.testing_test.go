package main

import (
	"fmt"
	"testing"
)

func add(x, y int) int {
	return x + y
}

func TestAdd(t *testing.T) {
	result := add(2, 3)
	if result != 5 {
		t.Errorf("Expected result to be 5, but got %v", result)
	}
}

func TestAddParameterized(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{4, 9, 13},
		{27, 41, 68},
		{15749, 516579, 532328},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("Adding %d and %d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			result := add(tt.a, tt.b)
			if result != tt.want {
				t.Errorf("Expected result to be %d, want %but got", tt.want, result)
			}
		})
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(1, 2)
	}
}
