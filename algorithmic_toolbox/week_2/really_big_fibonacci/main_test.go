package main_test

import (
	"testing"

	fib "github.com/tcc-sejohnson/ds-and-a/algorithmic_toolbox/week_2/really_big_fibonacci"
)

func TestPisanoPeriodLength(t *testing.T) {
	tests := []struct {
		Input int
		Want  int
	}{
		{Input: 2, Want: 3},
		{Input: 3, Want: 8},
	}

	for _, test := range tests {
		got := fib.FindPisanoPeriodLength(test.Input)
		if test.Want != got {
			t.Errorf("fib.FindPisanoPeriodLength returned incorrect data. Want: %v, got: %v", test.Want, got)
		}
	}
}

func TestFibModulo(t *testing.T) {
	tests := []struct {
		InputN int
		InputM int
		Want   int
	}{
		{InputN: 239, InputM: 1000, Want: 161},
		{InputN: 2816213588, InputM: 239, Want: 151},
	}

	for _, test := range tests {
		got := fib.FibModulo(test.InputN, test.InputM)
		if test.Want != got {
			t.Errorf("fib.FibModulo returned incorrect data. Want: %v, got: %v", test.Want, got)
		}
	}
}
