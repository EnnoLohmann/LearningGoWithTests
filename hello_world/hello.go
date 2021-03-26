package main

import (
	"fmt"
)

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const spanish = "Spanish"
const english = "English"
const french = "French"

func Hello(name, language string) string {
	helloPrefix := getHelloPrefix(language)

	if len(name) == 0 {
		name = "world"
	}

	return helloPrefix + name
}

func getHelloPrefix(language string) (prefix string) {
	switch language {
	case english:
		prefix = englishHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("", ""))
}
