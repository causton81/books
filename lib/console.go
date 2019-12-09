package lib

import (
	"bufio"
	"fmt"
	"github.com/causton81/books/boundary"
	"github.com/causton81/books/context"
	"os"
	"strconv"
	"strings"
)

func NewTextConsole() boundary.Console {
	cons := new(TextConsole)
	cons.r = bufio.NewReader(os.Stdin)
	return cons
}

type TextConsole struct {
	r *bufio.Reader
}

func (t *TextConsole) ReadLine() string {
	line, err := t.r.ReadString('\n')
	Must(err)
	return strings.TrimSpace(line)
}

func (t *TextConsole) Printf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

type ConsoleApp interface {
	Run(boundary.Console) int
}

type consoleApp struct {
}

func (c consoleApp) Run(cons boundary.Console) int {
	cons.Printf("Welcome to the Books App!\n")
	cmd := ""
	res := &boundary.QueryBookResponse{}
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
			res, err = context.QB.Execute(&boundary.QueryBookRequest{Query: queryString})
			Must(err)

		case "v":
			res := context.VL.Execute()
			cons.Printf("Your reading list:\n")
			for i, bk := range res.Books {
				cons.Printf("%d - %q\n", i, bk.Title)
			}
			cons.Printf("\n\n")
		default:
			n, err := strconv.Atoi(cmd)
			if nil == err && n < len(res.Books) {
				//cons.Printf("TODO: u picked %d %q\n", n, res.Books[n].Id)
				err = context.AB.Execute(boundary.AddBookRequest{Id: res.Books[n].Id})
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

func NewConsoleApp() ConsoleApp {
	return &consoleApp{}
}
