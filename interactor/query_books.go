package interactor

import (
	"github.com/causton81/books/boundary"
	"github.com/causton81/books/context"
)

func NewQueryBook() boundary.QueryBook {
	return &queryBooks{}
}

type queryBooks struct {
}

func (qb *queryBooks) Execute(req *boundary.QueryBookRequest) (*boundary.QueryBookResponse, error) {
	vols, err := context.VolumeSrv.FullTextSearch(req.Query)
	if nil != err {
		return nil, err
	}

	res := &boundary.QueryBookResponse{}

	for _, v := range vols {
		bk := context.BookGw.NewBook()
		bk.SetId(v.Id())
		bk.SetTitle(v.Title())
		context.BookGw.SaveBook(bk)

		res.Books = append(res.Books, &boundary.BookModel{
			Id:        v.Id(),
			Title:     v.Title(),
			Authors:   v.Authors(),
			Publisher: v.Publisher(),
		})
	}

	return res, nil
}

//func (q QueryBookRequest) RequestModel() {
//}
//
//type builder struct {
//	req QueryBookRequest
//}
//
//func (b builder) SetQuery(q string) boundary.RequestBuilder {
//	b.req.Query = q
//}
//
//func (b builder) Build() boundary.RequestModel {
//	return b.req
//}
