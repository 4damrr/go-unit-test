package main

import "testing"

func TestHitungFactorial(t *testing.T) {
	// setup
	tests := []struct {
		input    int
		expected int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{-1, -1},
	}

	// exercise
	for _, test := range tests {
		result := HitungFactorial(test.input)
		if result != test.expected {
			t.Errorf("SALAH! harusnya %v, bukan %v", test.expected, result)
		}
	}
}
