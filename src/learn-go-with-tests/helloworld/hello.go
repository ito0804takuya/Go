package main

import "fmt"

const japanese = "Japanese"
const french = "French"

const englishHelloPrefix = "Hello, "
const japaneseHelloPrefix = "こんにちは, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

// 返り値の定義でprefix変数を定義できる！ (Named return value)
// それで、それをそのまま関数内で使えるのはもちろん、return とするだけで、prefix変数が返される！
func greetingPrefix(language string) (prefix string) {
	switch language {
	case japanese:
		prefix = japaneseHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("John", ""))
}
