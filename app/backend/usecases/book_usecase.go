package usecases

import (
	"slices"
	"strings"

	"github.com/ttttai/golang/domain/entities"
	"github.com/ttttai/golang/domain/services"
)

type IBookUsecase interface {
	SearchBooks(title string, maxNum int, offset int) (*[]entities.BookInfo, error)
	GetBookInfoByBookId(id int) (*entities.BookInfo, error)
	CreateBookInfo(bookInfo *entities.BookInfo) (*entities.BookInfo, error)
	UpdateBook(book *entities.Book) (*entities.Book, error)
	DeleteBook(id int) error
	UpdateBookStatus(id int, bookStatus int) (*entities.Book, error)
	GetBookInfo(title string, status []int) (*[]entities.BookInfo, error)
	GetGeminiResponse(prompt string) ([]entities.BookInfo, error)
}

type BookUsecase struct {
	bookService      services.IBookService
	authorService    services.IAuthorService
	subjectService   services.ISubjectService
	geminiApiService services.IGeminiApiService
}

func NewBookUsecase(bookService services.IBookService, authorService services.IAuthorService, subjectService services.ISubjectService, geminiApiService services.IGeminiApiService) IBookUsecase {
	return &BookUsecase{
		bookService:      bookService,
		authorService:    authorService,
		subjectService:   subjectService,
		geminiApiService: geminiApiService,
	}
}

func (bu *BookUsecase) SearchBooks(title string, maxNum int, offset int) (*[]entities.BookInfo, error) {
	var bookInfo []entities.BookInfo

	bookInfoFromApi, err := bu.bookService.GetBooksFromNdlApi(title, maxNum, offset)
	if err != nil {
		return nil, err
	}

	// すでにDBに存在している場合、除外
	var excludedBookInfo []entities.BookInfo
	var bookInfoISBNs []int
	for _, bookInfoItem := range *bookInfoFromApi {
		book, err := bu.bookService.GetBookByISBN(bookInfoItem.Book.ISBN)
		if err != nil {
			return nil, err
		}

		if book == nil {
			excludedBookInfo = append(excludedBookInfo, bookInfoItem)
		} else {
			// APIから同じISBNの本を取得した場合スキップ
			if slices.Contains(bookInfoISBNs, book.ISBN) {
				continue
			}
			bookInfoISBNs = append(bookInfoISBNs, book.ISBN)

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

func (bu *BookUsecase) UpdateBook(book *entities.Book) (*entities.Book, error) {
	result, err := bu.bookService.UpdateBook(book)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (bu *BookUsecase) DeleteBook(id int) error {
	err := bu.bookService.DeleteBook(id)
	if err != nil {
		return err
	}

	return nil
}

func (bu *BookUsecase) UpdateBookStatus(id int, bookStatus int) (*entities.Book, error) {
	result, err := bu.bookService.UpdateBookStatus(id, bookStatus)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (bu *BookUsecase) GetBookInfo(title string, status []int) (*[]entities.BookInfo, error) {
	result, err := bu.bookService.GetBookInfo(title, status)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (bu *BookUsecase) GetGeminiResponse(prompt string) ([]entities.BookInfo, error) {
	bookInfo, err := bu.bookService.GetBookInfo("", []int{services.BOOK_STATUS_PURCHASED, services.BOOK_STATUS_READING, services.BOOK_STATUS_READ_COMPLETED})
	if err != nil {
		return nil, err
	}

	var bookTitles []string
	for _, bookInfoItem := range *bookInfo {
		bookTitles = append(bookTitles, bookInfoItem.Book.TitleName)
	}

	recommendationRequest := "以下の私が読んできた本を参考にして、おすすめの本を5冊挙げ、そのタイトルのみを箇条書きで提示してください。ただし、すでに読んだ本は提示しないてください。\\" + strings.Join(bookTitles, ",")
	recommendationTitles, err := bu.geminiApiService.GetGeminiResponse(recommendationRequest)
	if err != nil {
		return nil, err
	}

	var recommendationBooks []entities.BookInfo
	for _, recommendationTitle := range recommendationTitles {
		recommendationBookCandidates, err := bu.SearchBooks(recommendationTitle, 10, 1)
		if err != nil {
			return nil, err
		}
		if len(*recommendationBookCandidates) != 0 {
			recommendationBooks = append(recommendationBooks, (*recommendationBookCandidates)[0])
		}
	}
	return recommendationBooks, err
}
