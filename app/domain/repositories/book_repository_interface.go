package repositories

import "github.com/ttttai/golang/domain/entities"

type IBookRepository interface {
	GetBooksFromNdlApi(title string, maxNum int) (*[]entities.BookInfo, error)
	CreateBooks(books *[]entities.Book) (*[]entities.Book, error)
	CreateBook(book *entities.Book) (*entities.Book, error)
	CreateBookAuthors(bookAuthors *[]entities.BookAuthor) (*[]entities.BookAuthor, error)
	CreateBookSubjects(bookSubjects *[]entities.BookSubject) (*[]entities.BookSubject, error)
	GetBookByTitle(title string) (*entities.Book, error)
}
