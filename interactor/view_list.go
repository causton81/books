package interactor

import (
	"github.com/causton81/books/boundary"
	"github.com/causton81/books/context"
)

type viewList struct {
}

func (v viewList) Execute() boundary.ViewListResponse {
	res := boundary.ViewListResponse{}
	books := context.ListGw.GetBooks()
	for _, b := range books {
		res.Books = append(res.Books, boundary.ListBook{Title: b.Title()})
	}
	return res
}

func NewViewList() boundary.ViewList {
	return new(viewList)
}
