package repositories

import (
	"encoding/xml"
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
				ISBN  string `xml:"http://purl.org/dc/terms/ identifier"`
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
					Date string `xml:"http://purl.org/dc/terms/ date"`
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

func (br *BookRepository) GetBooksFromNdlApi(title string, maxNum int) (*[]entities.Book, error) {
	var books []entities.Book

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

		isbnStr := strings.ReplaceAll(bibResource.ISBN, "-", "")
		isbn, _ := strconv.Atoi(isbnStr)

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
			TitleName:         record.Record.BibResource.Title.Description.TitleName,
			TitleKana:         record.Record.BibResource.Title.Description.TitleKana,
			Authors:           authors,
			PublisherName:     record.Record.BibResource.Publisher.Agent.Name,
			PublisherNameKana: record.Record.BibResource.Publisher.Agent.NameKana,
			PublishDate:       record.Record.BibResource.Publisher.Date,
			Subjects:          subjects,
			Price:             price,
		}

		books = append(books, book)
	}

	return &books, nil
}
