package main

import (
	"encoding/json"
	"os"
)

// Bookworm represents a bookworm with a name and a list of books.
type Bookworm struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

// Book represents a book with a title and an author.
type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

// getBookworms returns a list of bookworms.
func getBookworms(filePath string) []Bookworm {
	var bookworms []Bookworm

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&bookworms)

	if err != nil {
		panic(err)
	}

	return bookworms
}
