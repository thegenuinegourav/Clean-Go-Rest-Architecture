package Book

import (
	"github.com/jinzhu/gorm"
)

type IBookRepository interface{
	GetBookById(id int) (*Book, error)
	GetAllBooks() (*[]Book, error)
	CreateBook(book *Book) (*Book, error)
	UpdateBook(book *Book) (*Book, error)
	DeleteBook(book *Book) error
}

type BookRepository struct {
	DB *gorm.DB
}

func (bookRepository *BookRepository) GetBookById(id int) (*Book, error) {
	var book Book
	result := bookRepository.DB.First(&book,id)
	return &book, result.Error
}

func (bookRepository *BookRepository) GetAllBooks() (*[]Book, error) {
	var book []Book
	result := bookRepository.DB.Find(&book)
	return &book, result.Error
}

func (bookRepository *BookRepository) CreateBook(book *Book) (*Book, error) {
	result := bookRepository.DB.Create(book)
	return book, result.Error
}

func (bookRepository *BookRepository) UpdateBook(book *Book) (*Book, error) {
	result := bookRepository.DB.Save(book)
	return book, result.Error
}

func (bookRepository *BookRepository) DeleteBook(book *Book) error {
	result := bookRepository.DB.Delete(book)
	return result.Error
}











