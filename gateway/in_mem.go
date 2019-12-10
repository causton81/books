package gateway

import (
	"github.com/causton81/books"
	"github.com/causton81/books/interactor"
	"log"
)

type listGateway struct {
	books map[string]bool
}

func (g *listGateway) AddBook(b interactor.Book) {
	id := b.(*inMemBook).id
	g.books[id] = true
}

func (g *listGateway) GetBooks() []interactor.Book {
	var res []interactor.Book
	for id := range g.books {
		res = append(res, books.BookGateway.FindById(id))
	}
	return res
}

func NewInMemoryListGw() interactor.ListGateway {
	gw := new(listGateway)
	gw.books = make(map[string]bool)
	return gw
}

type bookGateway struct {
	books map[string]*inMemBook
}

type inMemBook struct {
	id    string
	title string
}

func (b *inMemBook) Id() string {
	return b.id
}

func (b *inMemBook) Title() string {
	return b.title
}

func (b *inMemBook) SetTitle(title string) {
	b.title = title
}

func (b *inMemBook) SetId(id string) {
	b.id = id
}

func (gw *bookGateway) NewBook() interactor.Book {
	return new(inMemBook)
}

func (gw *bookGateway) SaveBook(b interactor.Book) {
	book := b.(*inMemBook)
	id := book.id
	gw.books[id] = book
}

func (gw *bookGateway) FindById(id string) interactor.Book {
	if b, found := gw.books[id]; found {
		return b
	}

	log.Fatal("bad book id")
	return nil
}

func NewInMemoryBookGw() interactor.BookGateway {
	gw := new(bookGateway)
	gw.books = make(map[string]*inMemBook)
	return gw
}
