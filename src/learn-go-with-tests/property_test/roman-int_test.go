package property_test

import (
	"fmt"
	"testing"
)

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Input int
		Want string
	}{
		{1, "I"},
		{2, "II"},
		{3, "III"},
		{4, "IV"},
		{5, "V"},
		{6, "VI"},
		{7, "VII"},
		{8, "VIII"},
		{9, "IX"},
		{10, "X"},
		{40, "XL"},
		{47, "XLVII"},
		{49, "XLIX"},
		{50, "L"},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("%d gets converted to %q", test.Input, test.Want), func(t *testing.T) {
			output := ConvertToRoman(test.Input)
			if output != test.Want {
				t.Errorf("got %q, want %q", output, test.Want)
			}
		})
	}
}
