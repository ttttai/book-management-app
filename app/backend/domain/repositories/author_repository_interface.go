package repositories

import "github.com/ttttai/golang/domain/entities"

type IAuthorRepository interface {
	CreateAuthor(author *entities.Author) (*entities.Author, error)
	CreateAuthors(authors *[]entities.Author) (*[]entities.Author, error)
	GetAuthorsByName(name string) (*[]entities.Author, error)
}
