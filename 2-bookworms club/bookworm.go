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
// It reads the data from a JSON file specified by filePath.
// The JSON file should contain an array of bookworm objects.
func getBookworms(filePath string) ([]Bookworm, error) {
	var bookworms []Bookworm

	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	err = json.NewDecoder(file).Decode(&bookworms)

	if err != nil {
		return nil, err
	}

	return bookworms, nil
}
