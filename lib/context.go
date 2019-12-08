package lib

type Volume interface {
	Id() string
	Title() string
	Publisher() string
	Authors() []string
}

type VolumeService interface {
	FullTextSearch(query string) ([]Volume, error)
}

var VolumeSrv VolumeService
