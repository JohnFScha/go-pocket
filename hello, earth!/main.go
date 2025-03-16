package main

import "fmt"

func main() {
	greeting := greet()
	fmt.Println(greeting)
}

// Greet returns a greeting message.
// It is a simple function that returns a string.
func greet() string {
	return "Hello, Earth!"
}
