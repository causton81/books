package impl

import (
	"github.com/causton81/books"
	"github.com/causton81/books/interactor"
	"github.com/stretchr/testify/require"
	"testing"
)

type stubListGw struct {

}

func (s stubListGw) AddBook(b interactor.Book) {
	panic("implement me")
}

type stubBook struct {
	title string
}

func (s stubBook) SetId(id string) {
	panic("implement me")
}

func (s stubBook) SetTitle(title string) {
	panic("implement me")
}

func (s stubBook) Id() string {
	panic("implement me")
}

func (s stubBook) Title() string {
	return s.title
}

func (s stubListGw) GetBooks() []interactor.Book {
	return []interactor.Book{
		stubBook{"title 1"},
		stubBook{"title 2"},
	}
}

func TestViewList(t *testing.T) {
	books.ListGateway = stubListGw{}
	vl := NewViewList()
	list := vl.Execute()
	a := require.New(t)
	a.Len(list.Books, 2)
}

