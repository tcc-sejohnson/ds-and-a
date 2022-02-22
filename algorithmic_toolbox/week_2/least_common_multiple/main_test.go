package main_test

import (
	"testing"

	lcm "github.com/tcc-sejohnson/ds-and-a/algorithmic_toolbox/week_2/least_common_multiple"
)

func TestLCM(t *testing.T) {
	tests := []struct {
		InputA int
		InputB int
		Want   int
	}{
		{InputA: 6, InputB: 8, Want: 24},
		{InputA: 761457, InputB: 614573, Want: 467970912861},
	}

	for _, test := range tests {
		got := lcm.LCM(test.InputA, test.InputB)
		if test.Want != got {
			t.Errorf("gcd.GCD returned incorrect results. Want: %v, got: %v", test.Want, got)
		}
	}
}
