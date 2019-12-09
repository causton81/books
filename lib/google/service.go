package google

import (
	"encoding/json"
	"fmt"
	"github.com/causton81/books/boundary"
	"github.com/causton81/books/lib"
	"net/http"
	"net/url"
)

func NewGoogleVolumeService() boundary.VolumeService {
	return new(googleVolumeService)
}

type googleVolumeService struct{}

type volume struct {
	*volumeResult
}

func (v volume) Id() string {
	return v.volumeResult.Id
}

func (v volume) Title() string {
	return v.volumeResult.VolumeInfo.Title
}

func (v volume) Publisher() string {
	return v.volumeResult.VolumeInfo.Publisher
}

func (v volume) Authors() []string {
	return v.volumeResult.VolumeInfo.Authors
}

type volumeResult struct {
	Id string
	// Each item in the list should include the book's author, title, and publishing company.
	VolumeInfo struct {
		Title     string
		Authors   []string
		Publisher string
	}
}

func (g *googleVolumeService) FullTextSearch(query string) ([]boundary.Volume, error) {
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
		Items      []*volumeResult
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

	var volumes []boundary.Volume
	for _, item := range out.Items {
		volumes = append(volumes, volume{item})
	}

	return volumes, nil
}
