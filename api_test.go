package books

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/url"
	"testing"
)

func TestGoogleBooksApi(t *testing.T) {
	// https://developers.google.com/books/docs/v1/reference/volumes/list.html
	// https://developers.google.com/books/docs/v1/reference/volumes.html#resource
	a := require.New(t)
	queryUrl, err := url.Parse("https://www.googleapis.com/books/v1/volumes")
	Must(err)
	queryValues := queryUrl.Query()
	queryValues.Set("q", "stormlight")
	queryValues.Set("maxResults", "5")
	queryValues.Set("projection", "lite") //  Includes a subset of fields in volumeInfo and accessInfo.
	queryValues.Set("printType", "books")
	queryValues.Set("langRestrict", "en")
	queryUrl.RawQuery = queryValues.Encode()
	res, err := http.Get(queryUrl.String())
	Must(err)
	defer res.Body.Close()
	t.Logf("status %d", res.StatusCode)
	a.Equal(http.StatusOK, res.StatusCode)

	dec := json.NewDecoder(res.Body)
	//var out map[string]interface{}
	var out struct {
		Kind  string
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
	Must(dec.Decode(&out))
	t.Logf("out %v", out)
	a.Equal("books#volumes", out.Kind)
}
