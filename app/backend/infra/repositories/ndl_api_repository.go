package repositories

import (
	"encoding/xml"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/ttttai/golang/domain/entities"
	"github.com/ttttai/golang/domain/repositories"
)

const NDL_SEARCH_API_URL = "https://ndlsearch.ndl.go.jp/api/sru"
const ISBN_URL = "http://ndl.go.jp/dcndl/terms/ISBN"
const OLD_ISBN_LEN = 10
const NEW_ISBN_LEN = 13

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

type NdlApiRepository struct{}

func NewNdlApiRepository() repositories.INdlApiRepository {
	return &NdlApiRepository{}
}

func (br *NdlApiRepository) GetBooksFromNdlApi(title string, maxNum int) (*[]entities.BookInfo, error) {
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
				if len(isbnStr) != NEW_ISBN_LEN {
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
				Name:     formatAuthorName(author.Author.Name),
				NameKana: formatAuthorName(author.Author.NameKana),
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

func formatAuthorName(input string) string {
	input = strings.ReplaceAll(input, ",", " ")

	re := regexp.MustCompile(`\d+|pub|-`)
	input = re.ReplaceAllString(input, "")

	input = strings.Join(strings.Fields(input), " ")

	return input
}
