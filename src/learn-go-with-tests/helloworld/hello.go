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

	prefix := englishHelloPrefix
	switch language {
	case japanese:
		prefix = japaneseHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	}

	return prefix + name
}
}

func main() {
	fmt.Println(Hello("John", ""))
}
