package property_test

import "strings"

type RomanNumeral struct {
	Value int
	Symbol string
}

var allRomanNumerals = []RomanNumeral {
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic int) string {
	var result strings.Builder

	for _, num := range allRomanNumerals {
		for arabic >= num.Value {
			result.WriteString(num.Symbol)
			arabic -= num.Value
		}
	}

	return result.String()
}
