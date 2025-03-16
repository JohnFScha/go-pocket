package main

import "fmt"

func main() {
	greeting := greet("en")
	fmt.Println(greeting)
}

// Package main provides a simple greeting function.
type language string

// phrasebook is a map of language codes to greeting messages.
var phrasebook = map[language]string{
	"en":  "Hello, Earth!",
	"es":  "¡Hola, Tierra!",
	"fr":  "Bonjour, Terre!",
	"gr":  "Γειά σου, Γη!",
	"he":  "שלום, ארץ!",
	"it":  "Ciao, Terra!",
	"akk": "𒀭𒊏𒁲𒀭𒊏𒁲",
	"":    "Language not supported",
}

// Greet returns a greeting message.
// It is a simple function that returns a string.
func greet(l language) string {
	greeting, ok := phrasebook[l]
	if !ok {
		return fmt.Sprintf("Language %q not supported", l)
	}
	return greeting
}
