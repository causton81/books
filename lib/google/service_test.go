package google

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGoogleBooksApi(t *testing.T) {
	//t.Skip("skipping exploratory test")
	a := require.New(t)
	srv := NewGoogleVolumeService()
	volumes, err := srv.FullTextSearch("stormlight")
	a.NoError(err)
	t.Logf("%v", volumes)
}
