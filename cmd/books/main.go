package main

import (
	"github.com/causton81/books"
	"github.com/causton81/books/console"
	"github.com/causton81/books/gateway"
	"github.com/causton81/books/interactor/impl"
	"github.com/causton81/books/lib/google"
	"os"
)

func main() {
	books.VolumeService = google.NewGoogleVolumeService()
	books.BookGateway = gateway.NewInMemoryBookGw()
	books.ListGateway = gateway.NewInMemoryListGw()
	books.QueryBook = impl.NewQueryBook()
	books.AddBookToList = impl.NewAddBookToList()
	books.ViewList = impl.NewViewList()
	os.Exit(console.NewConsoleApp().Run(console.NewTextConsole()))
}
