package repositories

import "github.com/ttttai/golang/domain/entities"

type INdlApiRepository interface {
	GetBooksFromNdlApi(title string, maxNum int, offset int) (*[]entities.BookInfo, error)
}
