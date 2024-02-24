package property_test

import "strings"

func ConvertToRoman(arabic int) string {
	// 文字列を構築するのに効率的なstrings.Builderを使う
	var result strings.Builder

	for arabic > 0 {
		switch {
		case arabic >= 10:
			result.WriteString("X")
			arabic -= 10
		case arabic == 9:
			result.WriteString("IX")
			arabic -= 9
		case arabic >= 5:
			result.WriteString("V")
			arabic -= 5
		case arabic == 4:
			result.WriteString("IV")
			arabic -= 4
		default:
			result.WriteString("I")
			arabic--
		}
	}

	// バッファを文字列におこして返す
	return result.String()
}
