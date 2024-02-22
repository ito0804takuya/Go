package property_test

import "strings"

func ConvertToRoman(arabic int) string {
	// 文字列を構築するのに効率的なstrings.Builderを使う
	var result strings.Builder

	for i := arabic; i > 0; i-- {
		if i == 4 {
			result.WriteString("IV")
			break
		}
		// バッファに溜め込む
		result.WriteString("I")
	}

	// バッファを文字列におこして返す
	return result.String()
}
