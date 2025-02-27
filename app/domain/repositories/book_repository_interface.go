package repositories

import "github.com/ttttai/golang/domain/entities"

type IBookRepository interface {
	CreateBooks(books *[]entities.Book) (*[]entities.Book, error)
	CreateBook(book *entities.Book) (*entities.Book, error)
	CreateBookAuthors(bookAuthors *[]entities.BookAuthor) (*[]entities.BookAuthor, error)
	CreateBookSubjects(bookSubjects *[]entities.BookSubject) (*[]entities.BookSubject, error)
	GetBooksByTitle(title string) (*[]entities.Book, error)
	GetBooksByFuzzyTitle(title string) (*[]entities.Book, error)
	GetBookByISBN(isbn int) (*entities.Book, error)
	GetBookInfoByISBNs(isbnSlices []int) (*[]entities.BookInfo, error)
	GetBookInfoByBookIds(ids []int) (*[]entities.BookInfo, error)
	DeleteBook(id int) error
}
