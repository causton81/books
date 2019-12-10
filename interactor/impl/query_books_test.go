package impl

import (
	"github.com/causton81/books"
	"github.com/causton81/books/interactor"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestQueryBooks(t *testing.T) {
	books.VolumeService = &bookServiceStub{}
	qb := NewQueryBook()
	res, err := qb.Execute(&interactor.QueryBookRequest{Query: "stormlight"})
	a := require.New(t)
	a.NoError(err)
	a.Len(res.Books, 2)
	assertBookModel(a, res.Books[0], "1")
	assertBookModel(a, res.Books[1], "2")
}

type bookServiceStub struct {
}

type volumeStub struct {
	id, title, publisher string
	authors              []string
}

func (v volumeStub) Id() string {
	return v.id
}

func (v volumeStub) Title() string {
	return v.title
}

func (v volumeStub) Publisher() string {
	return v.publisher
}

func (v volumeStub) Authors() []string {
	return v.authors
}

func (b bookServiceStub) FullTextSearch(query string) ([]interactor.Volume, error) {
	return []interactor.Volume{
		volumeStub{id: "unit-book-id-1", title: "unit title 1", authors: []string{"unit author 1"}, publisher: "unit publisher 1"},
		volumeStub{id: "unit-book-id-2", title: "unit title 2", authors: []string{"unit author 2"}, publisher: "unit publisher 2"},
	}, nil
}

func assertBookModel(a *require.Assertions, obtained *interactor.BookModel, suff string) {
	a.Equal(
		&interactor.BookModel{
			Id:        "unit-book-id-" + suff,
			Title:     "unit title " + suff,
			Authors:   []string{"unit author " + suff},
			Publisher: "unit publisher " + suff,
		},
		obtained,
	)
}
