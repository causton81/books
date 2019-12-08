package interactor

import (
	"github.com/causton81/books/lib"
)

type BookModel struct {
	Id        string
	Title     string
	Authors   []string
	Publisher string
}

type QueryBookResponse struct {
	Books []*BookModel
}

type QueryBookRequest struct {
	Query string
}

func NewQueryBook() *queryBooks {
	return &queryBooks{}
}

type queryBooks struct {
}

func (qb *queryBooks) Execute(req *QueryBookRequest) (*QueryBookResponse, error) {
	vols, err := lib.VolumeSrv.FullTextSearch(req.Query)
	if nil != err {
		return nil, err
	}

	res := &QueryBookResponse{}

	for _, v := range vols {
		res.Books = append(res.Books, &BookModel{
			Id:        v.Id(),
			Title:     v.Title(),
			Authors:   v.Authors(),
			Publisher: v.Publisher(),
		})
	}

	return res, nil
}
