package impl

import (
	"fmt"
	"github.com/causton81/books"
	"github.com/causton81/books/interactor"
	"strings"
)

type addBookToList struct {
}

func (a addBookToList) Execute(req interactor.AddBookRequest) error {
	id := strings.TrimSpace(req.Id)
	if "" == id {
		return fmt.Errorf("validate id: is blank")
	}
	book := books.BookGateway.FindById(id)
	books.ListGateway.AddBook(book)
	return nil
}

func NewAddBookToList() interactor.AddBookToList {
	return &addBookToList{}
}
