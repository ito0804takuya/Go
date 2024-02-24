package property_test

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumerals = RomanNumerals{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
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

type RomanNumerals []RomanNumeral

// ローマ数字を返す
func (r RomanNumerals) ValueOf(symbol string) int {
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}
	return 0
}

func ConvertToArabic(roman string) int {
	total := 0

	for i := 0; i < len(roman); i++ {
		symbol := roman[i]

		if i+1 < len(roman) && symbol == 'I' {
			nextSymbol := roman[i+1]

			potentialNumber := string([]byte{symbol, nextSymbol})

			value := allRomanNumerals.ValueOf(potentialNumber)

			if value != 0 {
				total += value
				// "IV"のように2文字を1つとして計算を行った場合は、iを1つ進める
				i++
			} else {
				total++
			}
		} else {
			total++
		}
	}

	return total
}