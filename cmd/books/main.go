package main

import (
	"github.com/causton81/books/context"
	"github.com/causton81/books/gateway"
	"github.com/causton81/books/interactor"
	"github.com/causton81/books/lib"
	"github.com/causton81/books/lib/google"
	"os"
)

func main() {
	context.VolumeSrv = google.NewGoogleVolumeService()
	context.BookGw = gateway.NewInMemoryBookGw()
	context.ListGw = gateway.NewInMemoryListGw()
	context.QB = interactor.NewQueryBook()
	context.AB = interactor.NewAddBookToList()
	context.VL = interactor.NewViewList()
	os.Exit(lib.NewConsoleApp().Run(lib.NewTextConsole()))
}
