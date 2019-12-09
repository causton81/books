package interactor

import (
	"github.com/causton81/books/boundary"
	"github.com/causton81/books/context"
	"github.com/causton81/books/gateway"
	"github.com/stretchr/testify/require"
	"testing"
)

type listGatewaySpy struct {
	lastId string
}

func (l *listGatewaySpy) AddBook(b boundary.Book) {
	l.lastId = b.Id()
}

func (l *listGatewaySpy) GetBooks() []boundary.Book {
	panic("implement me")
}

func TestAddBookToList(t *testing.T) {
	abtl := NewAddBookToList()
	bg := gateway.NewInMemoryBookGw()
	context.BookGw = bg
	b := bg.NewBook()
	b.SetId("42")
	bg.SaveBook(b)
	spy := &listGatewaySpy{}
	//lg := gateway.NewInMemoryListGw()
	context.ListGw = spy
	err := abtl.Execute(boundary.AddBookRequest{Id: "42"})

	a := require.New(t)
	a.NoError(err)
	a.Equal("42", spy.lastId)
}
