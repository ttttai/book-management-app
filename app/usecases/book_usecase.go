package usecases

import (
	"github.com/ttttai/golang/domain/entities"
	"github.com/ttttai/golang/domain/services"
)

type IBookUsecase interface {
	SearchBooks(title string, maxNum int) (*[]entities.BookInfo, error)
	GetBookInfoByBookId(id int) (*entities.BookInfo, error)
	CreateBookInfo(bookInfo *entities.BookInfo) (*entities.BookInfo, error)
}

type BookUsecase struct {
	bookService    services.IBookService
	authorService  services.IAuthorService
	subjectService services.ISubjectService
}

func NewBookUsecase(bookService services.IBookService, authorService services.IAuthorService, subjectService services.ISubjectService) IBookUsecase {
	return &BookUsecase{
		bookService:    bookService,
		authorService:  authorService,
		subjectService: subjectService,
	}
}

func (bu *BookUsecase) SearchBooks(title string, maxNum int) (*[]entities.BookInfo, error) {
	var bookInfo []entities.BookInfo

	booksFromDB, err := bu.bookService.GetBooksByFuzzyTitle(title)
	if err != nil {
		return nil, err
	}

	if len(*booksFromDB) >= maxNum {
		// DBにすでにmaxNum個存在している場合、DBから取り出す
		for i := 0; i < maxNum; i++ {
			bookInfoItem, err := bu.GetBookInfoByBookId((*booksFromDB)[i].ID)
			if err != nil {
				return nil, err
			}
			bookInfo = append(bookInfo, *bookInfoItem)
		}
	} else {
		bookInfoFromApi, err := bu.bookService.GetBooksFromNdlApi(title, maxNum)
		if err != nil {
			return nil, err
		}

		// すでにDBに存在している場合、除外
		var excludedBookInfo []entities.BookInfo
		for _, bookInfoItem := range *bookInfoFromApi {
			book, err := bu.bookService.GetBookByISBN(bookInfoItem.Book.ISBN)
			if err != nil {
				return nil, err
			}

			if book == nil {
				excludedBookInfo = append(excludedBookInfo, bookInfoItem)
			} else {
				bookInfoItem, err := bu.GetBookInfoByBookId(book.ID)
				if err != nil {
					return nil, err
				}
				bookInfo = append(bookInfo, *bookInfoItem)
			}
		}

		// DBに登録
		for _, excludedBookInfoItem := range excludedBookInfo {
			bookInfoItem, err := bu.CreateBookInfo(&excludedBookInfoItem)
			if err != nil {
				return nil, err
			}
			bookInfo = append(bookInfo, *bookInfoItem)
		}
	}

	return &bookInfo, nil
}

func (bu *BookUsecase) GetBookInfoByBookId(id int) (*entities.BookInfo, error) {
	book, err := bu.bookService.GetBookInfoByBookIds([]int{id})
	if err != nil {
		return nil, err
	}

	return &(*book)[0], nil
}

func (bu *BookUsecase) CreateBookInfo(bookInfo *entities.BookInfo) (*entities.BookInfo, error) {
	newBook, err := bu.bookService.CreateBook(&bookInfo.Book)
	if err != nil {
		return nil, err
	}

	bookAuthorRelations, err := bu.authorService.GetBookAuthorRelations(newBook, &bookInfo.Authors)
	if err != nil {
		return nil, err
	}
	_, err = bu.bookService.CreateBookAuthors(bookAuthorRelations)
	if err != nil {
		return nil, err
	}

	bookSubjectRelations, err := bu.subjectService.GetBookSubjectRelations(newBook, &bookInfo.Subjects)
	if err != nil {
		return nil, err
	}
	_, err = bu.bookService.CreateBookSubjects(bookSubjectRelations)
	if err != nil {
		return nil, err
	}

	bookInfoResponse, err := bu.bookService.GetBookInfoByBookIds([]int{newBook.ID})
	if err != nil {
		return nil, err
	}

	return &(*bookInfoResponse)[0], nil
}
