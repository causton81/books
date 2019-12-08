package books

import (
	"encoding/json"
	"fmt"
	"github.com/causton81/books/interactor"
	"github.com/causton81/books/lib"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/url"
	"testing"
)

type googleVolumeService struct{}

func (g *googleVolumeService) FullTextSearch(query string) ([]lib.Volume, error) {
	// https://developers.google.com/books/docs/v1/reference/volumes/list.html
	queryUrl, err := url.Parse("https://www.googleapis.com/books/v1/volumes")
	lib.Must(err)
	queryValues := queryUrl.Query()
	queryValues.Set("q", query)
	queryValues.Set("maxResults", "5")
	queryValues.Set("projection", "lite") //  Includes a subset of fields in volumeInfo and accessInfo.
	queryValues.Set("printType", "books")
	queryValues.Set("langRestrict", "en")
	queryUrl.RawQuery = queryValues.Encode()
	res, err := http.Get(queryUrl.String())
	if nil != err {
		return nil, err
	}
	defer res.Body.Close()

	dec := json.NewDecoder(res.Body)

	var out struct {
		Kind string
		// https://developers.google.com/books/docs/v1/reference/volumes.html#resource
		Items []struct {
			Id string
			// Each item in the list should include the book's author, title, and publishing company.
			VolumeInfo struct {
				Title     string
				Authors   []string
				Publisher string
			}
		}
		TotalItems int
	}
	err = dec.Decode(&out)
	if nil != err {
		return nil, err
	}
	//t.Logf("out %v", out)
	const BooksKind = "books#volumes"
	if BooksKind != out.Kind {
		return nil, fmt.Errorf("expecting response kind %q: got %q", BooksKind, out.Kind)
	}

	var volumes []lib.Volume
	for _, item := range out.Items {
		volumes = append(volumes, volumeStub{
			id:        item.Id,
			title:     item.VolumeInfo.Title,
			publisher: item.VolumeInfo.Publisher,
			authors:   item.VolumeInfo.Authors,
		})
	}

	return volumes, nil
}

func TestGoogleBooksApi(t *testing.T) {
	//t.Skip("skipping exploratory test")
	a := require.New(t)
	srv := NewGoogleVolumeService()
	volumes, err := srv.FullTextSearch("stormlight")
	a.NoError(err)
	t.Logf("%v", volumes)
}

func NewGoogleVolumeService() lib.VolumeService {
	return new(googleVolumeService)
}

func TestQueryBooks(t *testing.T) {
	lib.VolumeSrv = &bookServiceStub{}
	qb := interactor.NewQueryBook()
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

func (b bookServiceStub) FullTextSearch(query string) ([]lib.Volume, error) {
	return []lib.Volume{
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
