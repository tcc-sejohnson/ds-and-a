package main_test

import (
	"testing"

	fib "github.com/tcc-sejohnson/ds-and-a/algorithmic_toolbox/week_2/fibonacci_last_digit"
)

func TestFibonacci(t *testing.T) {
	tests := []struct {
		Input int
		Want  int
	}{
		{Input: 10, Want: 5},
		{Input: 3, Want: 2},
		{Input: 331, Want: 9},
		{Input: 327305, Want: 5},
	}

	for _, test := range tests {
		got := fib.Fib(test.Input)
		if test.Want != got {
			t.Errorf("main.Fib returned an incorrect result. Want: %v, got: %v", test.Want, got)
		}
	}
}
