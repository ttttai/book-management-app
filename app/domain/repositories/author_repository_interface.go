package repositories

import "github.com/ttttai/golang/domain/entities"

type IAuthorRepository interface {
	CreateAuthors(authors *[]entities.Author) (*[]entities.Author, error)
}
