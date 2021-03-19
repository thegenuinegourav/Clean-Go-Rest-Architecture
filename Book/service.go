package Book

type IBookService interface{
	GetBookService(id int) (*Book, error)
	GetBooksService() (*[]Book, error)
	PostBookService(book *Book) (*Book, error)
	UpdateBookService(id int, book *Book) (*Book, error)
	DeleteBookService(id int) error
}

type BookService struct {
	IBookRepository
}

func (bookService *BookService) GetBookService(id int) (*Book, error) {
	return bookService.GetBookById(id)
}

func (bookService *BookService) GetBooksService() (*[]Book, error) {
	return bookService.GetAllBooks()
}

func (bookService *BookService) PostBookService(book *Book) (*Book, error) {
	return bookService.CreateBook(book)
}

func (bookService *BookService) UpdateBookService(id int, book *Book) (*Book, error) {
	res, err := bookService.GetBookById(id)
	if err != nil {
		return nil, err
	}
	res.Members = book.Members
	res.User = book.User
	return bookService.UpdateBook(res)
}

func (bookService *BookService) DeleteBookService(id int) error {
	res, err := bookService.GetBookById(id)
	if err != nil {
		return err
	}
	return bookService.DeleteBook(res)
}
