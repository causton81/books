package impl

import (
	books2 "github.com/causton81/books"
	"github.com/causton81/books/interactor"
)

type viewList struct {
}

func (v viewList) Execute() interactor.ViewListResponse {
	res := interactor.ViewListResponse{}
	books := books2.ListGateway.GetBooks()
	for _, b := range books {
		res.Books = append(res.Books, interactor.ListBook{Title: b.Title()})
	}
	return res
}

func NewViewList() interactor.ViewList {
	return new(viewList)
}
