package context

import "github.com/causton81/books/boundary"

var VolumeSrv boundary.VolumeService

var QB boundary.QueryBook
var AB boundary.AddBookToList
var VL boundary.ViewList

var ListGw boundary.ListGateway
var BookGw boundary.BookGateway
