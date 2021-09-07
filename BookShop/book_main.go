// https://www.geeksforgeeks.org/bookshop-management-system-using-file-handling/
// Learn about reading input using fmt, bufio, os streams
package main

import (
	"bufio"
	"fmt"
	"github.com/google/uuid"
	"strings"
	"os"
	"os/exec"
	"time"
)

var reader bufio.Reader

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func welcomeScreen() int {
	var ch string
	var rv int

	prevLogin = true

	for rv == 0 {
		clearScreen()

		fmt.Println("---------------------------------")
		fmt.Println("      WELCOME TO BOOK STORE      ")
		fmt.Println("---------------------------------")
		fmt.Println()
		fmt.Println("          Login    Exit          ")
		fmt.Println()

		fmt.Printf("Choice: ")
		// reader := bufio.NewReader(os.Stdin)
		fmt.Scanf("%s", &ch)
		ch = strings.ToLower(ch)
		switch ch {
			case "login": rv = 1
			case "exit": os.Exit(0)
		}
	}

	return rv
}

type book struct {
	ISBN string
	Name string
	Author string
}

func (b book) GenISBN() string {
	id := uuid.NewString()
	if id != "" {
		b.ISBN = id
	}
	return b.ISBN
}

type bookShop interface {
	ControlPanel()
	AddBook()
	ShowBooks()
	CheckBook()
	UpdateBook()
	DeleteBook()
}
type bookShopImpl struct {}

func (bsi bookShopImpl) ControlPanel(prevLogin bool) {
	if prevLogin == true {
		clearScreen()
	}

	fmt.Println("----------------------------------------")
	fmt.Println("                HOMEPAGE                ")
	fmt.Println("----------------------------------------")
	fmt.Println("         1. Add book                    ")
	fmt.Println("         2. Display books               ")
	fmt.Println("         3. Check particular book       ")
	fmt.Println("         4. Update book                 ")
	fmt.Println("         5. Delete book                 ")
	fmt.Println("         6. Back to Welcome screen      ")

}
func (bsi bookShopImpl) AddBook() {
	clearScreen()

	var b book
	reader := bufio.NewReader(os.Stdin)
	// flush the reader
	fmt.Printf("ISBN: %v\n", b.GenISBN())
	fmt.Printf("Book Name: ")
	b.Name, _ = reader.ReadString('\n')
	b.Name = strings.TrimSpace(b.Name)
	fmt.Printf("Name Read: %v", b.Name)
	time.Sleep(10 * time.Second)
//	fmt.Scanln(&b.Name)
	fmt.Printf("Author: ")
	fmt.Scanln(&b.Author)

	// Write the book object to file

	fmt.Println("!!!   Book added successfully   !!!")
	fmt.Printf("%#v\n", b)
	time.Sleep(5 * time.Second)
}
func (bsi bookShopImpl) ShowBooks() {
	clearScreen()

	// Display all books
	fmt.Println("ShowBooks()")
}
func (bsi bookShopImpl) CheckBook() {
	clearScreen()

	// Read book ISBN and search and display
	fmt.Println("CheckBook()")
}
func (bsi bookShopImpl) UpdateBook() {
	clearScreen()

	// Read updated details and update the book
	fmt.Println("UpdateBook()")
}
func (bsi bookShopImpl) DeleteBook() {
	clearScreen()

	// Search and delete the book
	fmt.Println("DeleteBook()")
}
var bsi bookShopImpl
var prevLogin bool
func initBookShop() {
	var chCP int
	ch := welcomeScreen()
	//time.Sleep(2 * time.Second)
//	for {
		switch ch {
			case 1:		for {
						bsi.ControlPanel(prevLogin)
						//time.Sleep(5 * time.Second)
						fmt.Printf("Enter choice: ")
						fmt.Scan(&chCP)
						switch chCP {
							case 1: bsi.AddBook()
								prevLogin = false
							case 2: bsi.ShowBooks()
								prevLogin = false
							case 3: bsi.CheckBook()
								prevLogin = false
							case 4: bsi.UpdateBook()
								prevLogin = false
							case 5: bsi.DeleteBook()
								prevLogin = false
							case 6: initBookShop()
							default: fmt.Println("Wrong choice..")
						}
					}
			case -1:	return
		}
//	}
}

func main() {
	initBookShop()
}
