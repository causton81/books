package console

import (
	"bufio"
	"fmt"
	"github.com/causton81/books"
	"github.com/causton81/books/interactor"
	"github.com/causton81/books/lib"
	"os"
	"strconv"
	"strings"
)

func NewTextConsole() Console {
	cons := new(TextConsole)
	cons.r = bufio.NewReader(os.Stdin)
	return cons
}

type TextConsole struct {
	r *bufio.Reader
}

func (t *TextConsole) ReadLine() string {
	line, err := t.r.ReadString('\n')
	lib.Must(err)
	return strings.TrimSpace(line)
}

func (t *TextConsole) Printf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

type App interface {
	Run(Console) int
}

type consoleApp struct {
}

func (c consoleApp) Run(cons Console) int {
	cons.Printf("Welcome to the Books App!\n")
	cmd := ""
	res := &interactor.QueryBookResponse{}
	var err error
	for "x" != cmd {
		if 0 < len(res.Books) {
			cons.Printf("Results from your latest query; you can enter the number to save it to your list:\n")
		}
		for i, book := range res.Books {
			cons.Printf("%d - %q by %s (%s)\n", i, book.Title, strings.Join(book.Authors, ", "), book.Publisher)
		}

		cons.Printf(`q - query
v - view reading list
x - exit
Enter an option > `)
		cmd = cons.ReadLine()
		switch cmd {
		case "x":
			break
		case "q":
			cons.Printf("Query: ")
			queryString := cons.ReadLine()
			//qb := interactor.NewQueryBook()
			res, err = books.QueryBook.Execute(&interactor.QueryBookRequest{Query: queryString})
			lib.Must(err)

		case "v":
			res := books.ViewList.Execute()
			cons.Printf("Your reading list:\n")
			for i, bk := range res.Books {
				cons.Printf("%d - %q\n", i, bk.Title)
			}
			cons.Printf("\n\n")
		default:
			n, err := strconv.Atoi(cmd)
			if nil == err && n < len(res.Books) {
				//cons.Printf("TODO: u picked %d %q\n", n, res.Books[n].Id)
				err = books.AddBookToList.Execute(interactor.AddBookRequest{Id: res.Books[n].Id})
				if nil != err {
					cons.Printf(err.Error())
				}
			} else {
				cons.Printf("Sorry, I did not understand that.\n")
			}
		}
	}

	cons.Printf("Bye!")
	return 0
}

func NewConsoleApp() App {
	return &consoleApp{}
}
