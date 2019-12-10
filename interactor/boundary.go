package interactor

type QueryBook interface {
	Execute(request *QueryBookRequest) (*QueryBookResponse, error)
}

type QueryBookRequest struct {
	Query string
}

type QueryBookResponse struct {
	Books []*BookModel
}

type BookModel struct {
	Id        string
	Title     string
	Authors   []string
	Publisher string
}

type AddBookToList interface {
	Execute(req AddBookRequest) error
}

type AddBookRequest struct {
	Id string
}

type AddBookResponse struct {
}

type ViewList interface {
	Execute() ViewListResponse
}

type ViewListResponse struct {
	Books []ListBook
}

type ListBook struct {
	Title string
}

type Book interface {
	SetId(id string)
	SetTitle(title string)
	Id() string
	Title() string
}

type ListGateway interface {
	AddBook(b Book)
	GetBooks() []Book
}

type BookGateway interface {
	NewBook() Book
	SaveBook(Book)
	FindById(id string) Book
}

type Volume interface {
	Id() string
	Title() string
	Publisher() string
	Authors() []string
}

type VolumeService interface {
	FullTextSearch(query string) ([]Volume, error)
}
