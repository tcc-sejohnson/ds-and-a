package main_test

import (
	"testing"

	fib "github.com/tcc-sejohnson/ds-and-a/algorithmic_toolbox/week_2/fibonacci"
)

func TestFibonacci(t *testing.T) {
	tests := []struct {
		Input int
		Want  int
	}{
		{Input: 10, Want: 55},
		{Input: 3, Want: 2},
	}

	for _, test := range tests {
		got := fib.Fib(test.Input)
		if test.Want != got {
			t.Errorf("main.Fib returned an incorrect result. Want: %v, got: %v", test.Want, got)
		}
	}
}
