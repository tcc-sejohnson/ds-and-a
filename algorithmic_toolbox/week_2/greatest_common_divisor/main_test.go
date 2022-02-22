package main_test

import (
	"testing"

	gcd "github.com/tcc-sejohnson/ds-and-a/algorithmic_toolbox/week_2/greatest_common_divisor"
)

func TestGCD(t *testing.T) {
	tests := []struct {
		InputA int
		InputB int
		Want   int
	}{
		{InputA: 18, InputB: 35, Want: 1},
		{InputA: 28851538, InputB: 1183019, Want: 17657},
	}

	for _, test := range tests {
		got := gcd.GCD(test.InputA, test.InputB)
		if test.Want != got {
			t.Errorf("gcd.GCD returned incorrect results. Want: %v, got: %v", test.Want, got)
		}
	}
}
