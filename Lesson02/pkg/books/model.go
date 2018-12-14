package books

type Book struct {
	ISBN   string
	Title  string
	Author string
}

type BookDatabase interface {
	GetBooks() ([]Book, error)
	Initialize() error
}
