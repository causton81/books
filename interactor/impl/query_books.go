package impl

import (
	"github.com/causton81/books"
	"github.com/causton81/books/interactor"
)

func NewQueryBook() interactor.QueryBook {
	return &queryBooks{}
}

type queryBooks struct {
}

func (qb *queryBooks) Execute(req *interactor.QueryBookRequest) (*interactor.QueryBookResponse, error) {
	vols, err := books.VolumeService.FullTextSearch(req.Query)
	if nil != err {
		return nil, err
	}

	res := &interactor.QueryBookResponse{}

	for _, v := range vols {
		bk := books.BookGateway.NewBook()
		bk.SetId(v.Id())
		bk.SetTitle(v.Title())
		books.BookGateway.SaveBook(bk)

		res.Books = append(res.Books, &interactor.BookModel{
			Id:        v.Id(),
			Title:     v.Title(),
			Authors:   v.Authors(),
			Publisher: v.Publisher(),
		})
	}

	return res, nil
}
