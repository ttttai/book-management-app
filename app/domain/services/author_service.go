package services

import (
	"github.com/ttttai/golang/domain/entities"
	"github.com/ttttai/golang/domain/repositories"
)

type IAuthorService interface {
	CreateAuthor(author *entities.Author) (*entities.Author, error)
	CreateAuthors(authors *[]entities.Author) (*[]entities.Author, error)
	GetAuthorByName(name string) (*entities.Author, error)
	GetBookAuthorRelations(book *entities.Book, authors *[]entities.Author) (*[]entities.BookAuthor, error)
}

type AuthorService struct {
	authorRepository repositories.IAuthorRepository
}

func NewAuthorService(authorRepository repositories.IAuthorRepository) IAuthorService {
	return &AuthorService{
		authorRepository: authorRepository,
	}
}

func (as *AuthorService) CreateAuthor(author *entities.Author) (*entities.Author, error) {
	result, err := as.authorRepository.CreateAuthor(author)
	if err != nil {
		return nil, err
	}

	return result, nil
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

func (as *AuthorService) GetAuthorByName(name string) (*entities.Author, error) {
	result, err := as.authorRepository.GetAuthorByName(name)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (as *AuthorService) GetBookAuthorRelations(book *entities.Book, authors *[]entities.Author) (*[]entities.BookAuthor, error) {
	var bookAuthorRelations []entities.BookAuthor
	for _, author := range *authors {
		if author.Name == "" {
			continue
		}

		var authorID int
		existingAuthor, _ := as.GetAuthorByName(author.Name)
		if existingAuthor == nil {
			newAuthor, err := as.CreateAuthor(&author)
			if err != nil {
				return nil, err
			}
			authorID = newAuthor.ID
		} else {
			authorID = existingAuthor.ID
		}

		bookAuthorRelations = append(
			bookAuthorRelations,
			entities.BookAuthor{
				BookID:   book.ID,
				AuthorID: authorID,
			},
		)
	}

	return &bookAuthorRelations, nil
}
