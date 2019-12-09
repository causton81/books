package gateway

import (
	"github.com/causton81/books/context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInMemoryGateway(t *testing.T) {
	bg := NewInMemoryBookGw()
	context.BookGw = bg
	book1 := bg.NewBook()
	book1.SetId("id1")
	book1.SetTitle("title 1")
	bg.SaveBook(book1)

	book2 := bg.NewBook()
	book2.SetId("id2")
	book2.SetTitle("title 2")
	bg.SaveBook(book2)

	lg := NewInMemoryListGw()
	lg.AddBook(book1)
	lg.AddBook(book2)

	books := lg.GetBooks()
	a := require.New(t)
	a.Len(books, 2)
	idx1 := 0
	idx2 := 1
	if "id2" == books[0].Id() {
		idx1, idx2 = 1, 0
	}

	a.Equal("id1", books[idx1].Id())
	a.Equal("title 1", books[idx1].Title())
	a.Equal("id2", books[idx2].Id())
	a.Equal("title 2", books[idx2].Title())
}
