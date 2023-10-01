package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const indonesianHelloPrefix = "Halo, "
const japaneseHelloPrefix = "Konnichiwa, "

func Hello(name, language string) string {
	if name == "" {
		return "Hello, World"
	}

	prefix := englishHelloPrefix

	switch language {
	case "French":
		prefix = frenchHelloPrefix
	case "Spanish":
		prefix = spanishHelloPrefix
	case "Indonesian":
		prefix = indonesianHelloPrefix
	case "Japanese":
		prefix = japaneseHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
