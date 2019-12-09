package interactor

import (
	"fmt"
	"github.com/causton81/books/boundary"
	"github.com/causton81/books/context"
	"strings"
)

type addBookToList struct {
}

func (a addBookToList) Execute(req boundary.AddBookRequest) error {
	id := strings.TrimSpace(req.Id)
	if "" == id {
		return fmt.Errorf("validate id: is blank")
	}
	book := context.BookGw.FindById(id)
	context.ListGw.AddBook(book)
	return nil
}

func NewAddBookToList() boundary.AddBookToList {
	return &addBookToList{}
}
