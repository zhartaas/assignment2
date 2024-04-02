package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var books []Book
var magazines []Magazine

func main() {
	librarian := Librarian{name: "admin", password: "admin123"}

	fmt.Println("Welcome to the Library Management System")

	// Loop for CLI interaction
	for {
		fmt.Println("\nSelect an option:")
		fmt.Println("1. Add Book")
		fmt.Println("2. Add Magazine")
		fmt.Println("3. List All Books")
		fmt.Println("4. List All Magazines")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		reader := bufio.NewReader(os.Stdin)
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)
		fmt.Println("\n\n\n\n\n")

		switch choice {
		case "1":
			librarian.AddBook()
		case "2":
			librarian.AddMagazine()
		case "3":
			listAllBooks()
		case "4":
			listAllMagazines()
		case "5":
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func listAllBooks() {
	fmt.Println("List of all books:")
	for i, book := range books {
		fmt.Printf("%d. %s by %s\n", i+1, book.Title, book.Author)
	}
}

func listAllMagazines() {
	fmt.Println("List of all magazines:")
	for i, magazine := range magazines {
		fmt.Printf("%d. %s (%s)\n", i+1, magazine.Title, magazine.Period)
	}
}

type User interface {
	Borrow(item LibraryItem)
}

type Librarian struct {
	name     string
	password string
}

func (l Librarian) AddBook() {
	var name string
	var author string
	fmt.Print("Enter name of the book: ")
	fmt.Scanf("%v\n", &name)

	fmt.Print("Enter author name of the book: ")
	fmt.Scanf("%v\n", &author)

	book := Book{Title: name, Author: author}

	books = append(books, book)
}

func (l Librarian) AddMagazine() {
	var title string
	var period string
	fmt.Print("Enter title of the magazine: ")
	fmt.Scanf("%v\n", &title)

	fmt.Print("Enter period of the magazine: ")
	fmt.Scanf("%v\n", &period)

	magazine := Magazine{Title: title, Period: period}

	magazines = append(magazines, magazine)
}

func (l Librarian) Borrow(item LibraryItem) {
	fmt.Println("Librarian borrowed item:")
	item.Display()
}

type Patron struct {
	name     string
	password string
}

func (p Patron) Borrow(item LibraryItem) {
	fmt.Println("Patron borrowed item:")
	item.Display()
}

type LibraryItem interface {
	Display()
}

type Book struct {
	Title  string
	Author string
}

func (b Book) Display() {
	fmt.Printf("Book: %s by %s\n", b.Title, b.Author)
}

type Magazine struct {
	Title  string
	Period string
}

func (m Magazine) Display() {
	fmt.Printf("Magazine: %s (%s)\n", m.Title, m.Period)
}
