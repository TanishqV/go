package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/google/uuid"
	"io"
	"strings"
	"os"
	"os/exec"
	"runtime"
)

const (
	FILENAME = "book_data.dat"
)

var reader bufio.Reader

func clearScreen() {
	if runtime.GOOS == "linux" {
		func() {
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
		}()
	} else if runtime.GOOS == "windows" {
		func() {
			cmd := exec.Command("cmd", "/c", "cls")
			cmd.Stdout = os.Stdout
			cmd.Run()
		}()
	} else {
		fmt.Printf("\nUnknown OS\n")
	}
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
		fmt.Scanf("%s", &ch)
		ch = strings.ToLower(ch)
		switch ch {
			case "login": rv = 1
			case "exit": os.Exit(0)
		}
	}

	return rv
}

func BytesToFixed(str string) (res [50]byte) {
	copy(res[:], str[:])
	return res
}

type book struct {
	ISBN [50]byte
	Name [50]byte
	Author [50]byte
}

func (b *book) GenISBN() string {
	id := uuid.NewString()
	if id != "" {
		(*b).ISBN = BytesToFixed(id)
	}
	return id
}

func (b *book) Display() {
	fmt.Printf("ISBN: %s\nName: %s\nAuthor: %s\n", b.ISBN, b.Name, b.Author)
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
	fmt.Printf("ISBN: %v\n", b.GenISBN())
	fmt.Printf("Book Name: ")
	b_name, _ := reader.ReadString('\n')
	b_name = strings.TrimSpace(b_name)
	fmt.Printf("Author: ")
	b_author, _ := reader.ReadString('\n')
	b_author = strings.TrimSpace(b_author)

	fp, err := os.OpenFile(FILENAME, os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0744)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	b.Name = BytesToFixed(b_name)
	b.Author = BytesToFixed(b_author)

	var buf bytes.Buffer
	binary.Write(&buf, binary.BigEndian, &b)
	_, err = fp.Write(buf.Bytes())
	if err != nil {
		panic(err)
	}

	fmt.Println("!!!   Book added successfully   !!!")
	b.Display()
	fmt.Println()
}
func (bsi bookShopImpl) ShowBooks() {
	// clearScreen()

	fp, err := os.OpenFile(FILENAME, os.O_RDONLY, 0744)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	var b book
	var count uint8 = 1
	for {
		rd_bytes := make([]byte, 150)
		_, err := fp.Read(rd_bytes)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		buf := bytes.NewBuffer(rd_bytes)
		err = binary.Read(buf, binary.BigEndian, &b)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Record #%d\n", count)
		b.Display()
		fmt.Println("-----------------------")
		count++
	}
}

func (bsi bookShopImpl) CheckBook() {
//	clearScreen()

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter ISBN or its prefix: ")
	isbn, _ := reader.ReadString('\n')
	isbn = strings.TrimSpace(isbn)

	resBook, _, ok := bsi.find(isbn)
	if ok == true {
		resBook.Display()
	} else {
		fmt.Printf("Book with ISBN: %s, NOT FOUND\n", isbn)
	}
}

func (bsi bookShopImpl) find(isbn string) (res book, offset int64, found bool) {

	fp, err := os.OpenFile(FILENAME, os.O_RDONLY, 0744)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	var b book
	for {
		rd_bytes := make([]byte, 150)
		_, err := fp.Read(rd_bytes)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		buf := bytes.NewBuffer(rd_bytes)
		err = binary.Read(buf, binary.BigEndian, &b)
		if err != nil {
			panic(err)
		}
		offset += 150
		if strings.HasPrefix(string(b.ISBN[:]), isbn) == true {
			found = true
			res = b
			break
		}
	}
	offset -=150
	return res, offset, found
}
func (bsi bookShopImpl) UpdateBook() {
	// clearScreen()

	// Read updated details and update the book
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter ISBN or its prefix: ")
	isbn , _ := reader.ReadString('\n')
	isbn = strings.TrimSpace(isbn)

	res, offset, ok := bsi.find(isbn)
	if ok == true {
		res.Display()
		fmt.Println("Enter new details for the same ISBN")
		fmt.Print("Name: ")
		b_name, _ := reader.ReadString('\n')
		b_name = strings.TrimSpace(b_name)
		fmt.Print("Author: ")
		b_author, _ := reader.ReadString('\n')
		b_author = strings.TrimSpace(b_author)

		if string(res.Name[:]) == b_name && string(res.Author[:]) == b_author {
			fmt.Println("No change.. skipping rewrite\n")
			return
		} else {
			// Seek to the offset and overwrite the existing record with the new one
			res.Name = BytesToFixed(b_name)
			res.Author = BytesToFixed(b_author)
			fp, err := os.OpenFile(FILENAME, os.O_WRONLY, 0744)
			if err != nil {
				panic(err)
			}
			defer fp.Close()

			var buf bytes.Buffer
			binary.Write(&buf, binary.BigEndian, &res)
			// bytes_written = _
			_, err = fp.WriteAt(buf.Bytes(), offset)
			if err != nil {
				panic(err)
			}
			// fmt.Printf("This bytes_written: %d\n", bytes_written)

			// fmt.Println("Book details updated for ISBN: %s", res.ISBN)
		}
	} else {
		fmt.Printf("Book with ISBN: %s, NOTFOUND\n", isbn)
	}
}
func (bsi bookShopImpl) DeleteBook() {
	// clearScreen()

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter ISBN or its prefix: ")
	isbn, _ := reader.ReadString('\n')
	isbn = strings.TrimSpace(isbn)

	res, _, ok := bsi.find(isbn)
	if ok == true {
		fmt.Println("Following book will be removed:")
		res.Display()

		var ch string = " "
		for ch[0] != 'y' && ch[0] != 'Y' && ch[0] != 'n' && ch[0] != 'N' {
			fmt.Print("Are you sure(Y/N)?")
			ch, _ = reader.ReadString('\n')
		}
		if ch[0] == 'Y' || ch[0] == 'y' {
			fp_read, err := os.OpenFile(FILENAME, os.O_RDONLY, 0744)
			if err != nil {
				panic(err)
			}
			defer fp_read.Close()

			fp_write, err := os.OpenFile("tmp.dat", os.O_WRONLY | os.O_CREATE, 0744)
			if err != nil {
				panic(err)
			}
			defer fp_write.Close()

			var b book
			rd_bytes := make([]byte, 150)
			for {
				_, err := fp_read.Read(rd_bytes)
				if err != nil {
					if err == io.EOF {
						break
					}
					panic(err)
				}
				// Buffer is drained on each read from it
				buf := bytes.NewBuffer(rd_bytes)
				buff_copy := bytes.NewBuffer(buf.Bytes())
				err = binary.Read(buff_copy, binary.BigEndian, &b)
				if err != nil {
					panic(err)
				}
				fmt.Println("Book read:")
				b.Display()
				if b.ISBN == res.ISBN {
					continue
				}

				_, err = fp_write.Write(buf.Bytes())
				if err != nil {
					panic(err)
				}
			}

			err = os.Remove(FILENAME)
			if err != nil {
				panic(err)
			}

			err = os.Rename("tmp.dat", FILENAME)
			if err != nil {
				panic(err)
			}

			fmt.Println("Book deleted successfully")
		}
	} else {
		fmt.Printf("Book with ISBN: %s, NOT FOUND\n", isbn)
	}
}
var bsi bookShopImpl
var prevLogin bool
func initBookShop() {
	var chCP int
	ch := welcomeScreen()
	switch ch {
		case 1:		for {
					bsi.ControlPanel(prevLogin)
					fmt.Printf("\nEnter choice: ")
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

}

func main() {
	initBookShop()
}
