package services

import (
	"github.com/ttttai/golang/domain/entities"
	"github.com/ttttai/golang/domain/repositories"
)

type IAuthorService interface {
	CreateAuthors(authors *[]entities.Author) (*[]entities.Author, error)
}

type AuthorService struct {
	authorRepository repositories.IAuthorRepository
}

func NewAuthorService(authorRepository repositories.IAuthorRepository) IAuthorService {
	return &AuthorService{
		authorRepository: authorRepository,
	}
}

func (as *AuthorService) CreateAuthors(authors *[]entities.Author) (*[]entities.Author, error) {
	if len(*authors) == 0 {
		return nil, nil
	}

	result, err := as.authorRepository.CreateAuthors(authors)
	if err != nil {
		return nil, err
	}

	return result, nil
}
