package iteration

import "strings"

const repeatCount = 5

func Repeat(character string) (repeated string) {
	return strings.Repeat(character, repeatCount)
}
