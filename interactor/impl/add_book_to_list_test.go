package impl

import (
	"github.com/causton81/books"
	"github.com/causton81/books/gateway"
	"github.com/causton81/books/interactor"
	"github.com/stretchr/testify/require"
	"testing"
)

type listGatewaySpy struct {
	lastId string
}

func (l *listGatewaySpy) AddBook(b interactor.Book) {
	l.lastId = b.Id()
}

func (l *listGatewaySpy) GetBooks() []interactor.Book {
	panic("implement me")
}

func TestAddBookToList(t *testing.T) {
	abtl := NewAddBookToList()
	bg := gateway.NewInMemoryBookGw()
	books.BookGateway = bg
	b := bg.NewBook()
	b.SetId("42")
	bg.SaveBook(b)
	spy := &listGatewaySpy{}
	//lg := gateway.NewInMemoryListGw()
	books.ListGateway = spy
	err := abtl.Execute(interactor.AddBookRequest{Id: "42"})

	a := require.New(t)
	a.NoError(err)
	a.Equal("42", spy.lastId)
}
