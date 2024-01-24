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

	if language == japanese {
		return japaneseHelloPrefix + name
	}
	if language == french {
		return frenchHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("John", ""))
}
