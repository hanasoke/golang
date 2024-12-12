package main

import "fmt"

type Book struct {
	ID        int
	Publisher string
	Writer    string
	IsActive  bool
}

func (book Book) display() string {
	return fmt.Sprintf("Publisher : %s, Writer : %s", book.Publisher, book.Writer)
}
func displayBook(book Book) string {
	return fmt.Sprintf("Publisher : %s, Writer : %s", book.Publisher, book.Writer)
}

type Total struct {
	Publisher   string
	Admin       Book
	Books       []Book
	IsAvailable bool
}

func main() {
	book := Book{1, "Manga", "Masashi Kishimoto", true}
	result := book.display()
	fmt.Println(result)

	book2 := Book{2, "Jambu", "Anna Mortgage", true}
	fmt.Println(book2.display())
}
