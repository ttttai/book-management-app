package repositories

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/ttttai/golang/domain/entities"
	"github.com/ttttai/golang/domain/repositories"
	"gorm.io/gorm"
)

const NDL_SEARCH_API_URL = "https://ndlsearch.ndl.go.jp/api/sru"
const ISBN_URL = "http://ndl.go.jp/dcndl/terms/ISBN"
const OLD_ISBN_LEN = 10
const NEW_ISBN_LEN = 13

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) repositories.IBookRepository {
	return &BookRepository{
		db: db,
	}
}

type XML struct {
	XMLName xml.Name `xml:"searchRetrieveResponse"`
	Records []struct {
		Record struct {
			BibResource struct {
				Identifiers []struct {
					Value    string `xml:",chardata"`
					Datatype string `xml:"http://www.w3.org/1999/02/22-rdf-syntax-ns# datatype,attr"`
				} `xml:"http://purl.org/dc/terms/ identifier"`
				Title struct {
					Description struct {
						TitleName string `xml:"http://www.w3.org/1999/02/22-rdf-syntax-ns# value"`
						TitleKana string `xml:"http://ndl.go.jp/dcndl/terms/ transcription"`
					} `xml:"http://www.w3.org/1999/02/22-rdf-syntax-ns# Description"`
				} `xml:"http://purl.org/dc/elements/1.1/ title"`
				Authors []struct {
					Author struct {
						Name     string `xml:"http://xmlns.com/foaf/0.1/ name"`
						NameKana string `xml:"http://ndl.go.jp/dcndl/terms/ transcription"`
					} `xml:"http://xmlns.com/foaf/0.1/ Agent"`
				} `xml:"http://purl.org/dc/terms/ creator"`
				Publisher struct {
					Agent struct {
						Name     string `xml:"http://xmlns.com/foaf/0.1/ name"`
						NameKana string `xml:"http://ndl.go.jp/dcndl/terms/ transcription"`
					} `xml:"http://xmlns.com/foaf/0.1/ Agent"`
					Date *string `xml:"http://purl.org/dc/terms/ date"`
				} `xml:"http://purl.org/dc/terms/ publisher"`
				Subjects []struct {
					Descriptions struct {
						SubjectName string `xml:"http://www.w3.org/1999/02/22-rdf-syntax-ns# value"`
						SubjectKana string `xml:"http://ndl.go.jp/dcndl/terms/ transcription"`
					} `xml:"http://www.w3.org/1999/02/22-rdf-syntax-ns# Description"`
				} `xml:"http://purl.org/dc/terms/ subject"`
				Price string `xml:"http://ndl.go.jp/dcndl/terms/ price"`
			} `xml:"http://ndl.go.jp/dcndl/terms/ BibResource"`
		} `xml:"http://www.w3.org/1999/02/22-rdf-syntax-ns# RDF"`
	} `xml:"records>record>recordData"`
}

func (br *BookRepository) GetBooksFromNdlApi(title string, maxNum int) (*[]entities.BookInfo, error) {
	var bookInfo []entities.BookInfo

	encodedTitle := url.PathEscape(title)
	res, err := http.Get(NDL_SEARCH_API_URL + "?operation=searchRetrieve" + "&recordPacking=xml" + "&recordSchema=dcndl" + "&maximumRecords=" + strconv.Itoa(maxNum) + "&query=title=" + encodedTitle)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var result XML
	if err := xml.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	for _, record := range result.Records {
		bibResource := record.Record.BibResource

		var isbn int
		existsIsbn := false
		for _, identifier := range bibResource.Identifiers {
			if identifier.Datatype == ISBN_URL {
				isbnValue := identifier.Value
				isbnStr := strings.ReplaceAll(isbnValue, "-", "")
				if len(isbnStr) != OLD_ISBN_LEN && len(isbnStr) != NEW_ISBN_LEN {
					continue
				}
				isbn, err = strconv.Atoi(isbnStr)
				if err != nil {
					return nil, err
				}
				existsIsbn = true
				break
			}
		}

		if !existsIsbn {
			continue
		}

		priceStr := strings.ReplaceAll(bibResource.Price, "円", "")
		price, _ := strconv.Atoi(priceStr)

		var authors []entities.Author
		for _, author := range bibResource.Authors {
			authors = append(authors, entities.Author{
				Name:     author.Author.Name,
				NameKana: author.Author.NameKana,
			})
		}

		var subjects []entities.Subject
		for _, subject := range bibResource.Subjects {
			if subject.Descriptions.SubjectName != "" {
				subjects = append(subjects, entities.Subject{
					SubjectName: subject.Descriptions.SubjectName,
					SubjectKana: subject.Descriptions.SubjectKana,
				})
			}
		}

		book := entities.Book{
			ISBN:              isbn,
			TitleName:         bibResource.Title.Description.TitleName,
			TitleNameKana:     bibResource.Title.Description.TitleKana,
			PublisherName:     bibResource.Publisher.Agent.Name,
			PublisherNameKana: bibResource.Publisher.Agent.NameKana,
			PublishDate:       bibResource.Publisher.Date,
			Price:             price,
		}

		bookInfoItem := entities.BookInfo{
			Book:     book,
			Authors:  authors,
			Subjects: subjects,
		}

		bookInfo = append(bookInfo, bookInfoItem)
	}

	return &bookInfo, nil
}

func (br *BookRepository) CreateBook(book *entities.Book) (*entities.Book, error) {
	result := br.db.Create(book)
	if result.Error != nil {
		return nil, result.Error
	}

	return book, nil
}

func (br *BookRepository) CreateBooks(books *[]entities.Book) (*[]entities.Book, error) {
	result := br.db.Create(books)
	if result.Error != nil {
		return nil, result.Error
	}

	return books, nil
}

func (br *BookRepository) CreateBookAuthors(bookAuthors *[]entities.BookAuthor) (*[]entities.BookAuthor, error) {
	result := br.db.Create(bookAuthors)
	if result.Error != nil {
		return nil, result.Error
	}

	return bookAuthors, nil
}

func (br *BookRepository) CreateBookSubjects(bookSubjects *[]entities.BookSubject) (*[]entities.BookSubject, error) {
	result := br.db.Create(bookSubjects)
	if result.Error != nil {
		return nil, result.Error
	}

	return bookSubjects, nil
}

func (br *BookRepository) GetBooksByTitle(title string) (*[]entities.Book, error) {
	var book []entities.Book

	result := br.db.Where("title_name = ?", title).Find(&book)
	if result.Error != nil {
		return nil, result.Error
	}

	return &book, nil
}

func (br *BookRepository) GetBookByISBN(isbn int) (*entities.Book, error) {
	var book entities.Book

	result := br.db.Where("isbn = ?", isbn).First(&book)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &book, nil
}

func (br *BookRepository) GetBookInfoByISBN(isbnSlices []int) (*[]entities.BookInfo, error) {
	var bookInfo []entities.BookInfo
	var book []entities.Book

	result := br.db.Preload("Authors").Preload("Subjects").Where("isbn IN ?", isbnSlices).Find(&book)
	if result.Error != nil {
		return nil, result.Error
	}
	fmt.Println(book)

	return &bookInfo, nil
}
