package main

import "fmt"

type Library struct {
	ID        int
	Title     string
	Publisher string
	Author    string
	Year      string
	IsActive  bool
}

func (library Library) Display() string {
	return fmt.Sprintf("Title : %s, Publisher : %s, Author : %s, Year : %s", library.Title, library.Publisher, library.Author, library.Year)
}

func displayLibrary(library Library) string {
	return fmt.Sprintf("Title : %s, Publisher : %s, Author : %s, Year : %s", library.Title, library.Publisher, library.Author, library.Year)
}

type Book struct {
	Title       string
	Admin       Library
	Users       []Library
	IsAvailable bool
}

func (book Book) DisplayBook() {
	fmt.Printf("Title : %s", book.Title)
	fmt.Println("")
	fmt.Printf("Member count : %d", len(book.Users))
	fmt.Println("")

	fmt.Println("Book title : ")
	for _, book := range book.Users {
		fmt.Println(book.Title)
	}
}

func displayBook(book Book) {
	fmt.Print("Title : %s", book.Title)
	fmt.Println("")
	fmt.Printf("Member count : %d", len(book.Users))
	fmt.Println("")

	fmt.Println("Book title :")
	for _, book := range book.Users {
		fmt.Println(book.Title)
	}
}

func main() {
	book := Library{1, "Laskar Pelangi", "Gramedia", "Andrea Hirata", "1999", true}
	result := book.Display()

	fmt.Println(result)

	book2 := Library{2, "Sebuah seni untuk bersikap bodo amat", "Gramedia", "Mark Manson", "2000", true}
	fmt.Println(book2.Display())

	books := []Library{book, book2}
	library := Book{"Ikigai", book, books, true}

	library.DisplayBook()
}
