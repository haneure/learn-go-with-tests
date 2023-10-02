package main

import "fmt"

const (
	french     = "French"
	spanish    = "Spanish"
	indonesian = "Indonesian"
	japanese   = "Japanese"

	englishHelloPrefix    = "Hello, "
	spanishHelloPrefix    = "Hola, "
	frenchHelloPrefix     = "Bonjour, "
	indonesianHelloPrefix = "Halo, "
	japaneseHelloPrefix   = "Konnichiwa, "
)

func Hello(name, language string) string {
	if name == "" {
		return "Hello, World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case indonesian:
		prefix = indonesianHelloPrefix
	case japanese:
		prefix = japaneseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
