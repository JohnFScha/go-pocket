package main

import (
	"fmt"
	"os"
)

func main() {
	bookworms, err := getBookworms("testdata/bookworms.json")

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load Bookworms: ")
		os.Exit(1)
	}

	for _, bookworm := range bookworms {
		println("Name:", bookworm.Name)
		for _, book := range bookworm.Books {
			println("  Title:", book.Title)
			println("  Author:", book.Author)
		}
	}
}
