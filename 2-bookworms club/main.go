package main

func main() {
	bookworms := getBookworms("testdata/bookworms.json")
	for _, bookworm := range bookworms {
		println("Name:", bookworm.Name)
		for _, book := range bookworm.Books {
			println("  Title:", book.Title)
			println("  Author:", book.Author)
		}
	}
}
