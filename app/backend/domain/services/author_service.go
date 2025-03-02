package services

import (
	"slices"

	"github.com/ttttai/golang/domain/entities"
	"github.com/ttttai/golang/domain/repositories"
)

type IAuthorService interface {
	CreateAuthor(author *entities.Author) (*entities.Author, error)
	CreateAuthors(authors *[]entities.Author) (*[]entities.Author, error)
	GetAuthorsByName(name string) (*[]entities.Author, error)
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

func (as *AuthorService) GetAuthorsByName(name string) (*[]entities.Author, error) {
	result, err := as.authorRepository.GetAuthorsByName(name)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (as *AuthorService) GetBookAuthorRelations(book *entities.Book, authors *[]entities.Author) (*[]entities.BookAuthor, error) {
	var bookAuthorRelations []entities.BookAuthor
	var authorIDs []int
	for _, author := range *authors {
		if author.Name == "" {
			continue
		}

		var authorID int
		existingAuthor, _ := as.GetAuthorsByName(author.Name)
		if len(*existingAuthor) == 0 {
			newAuthor, err := as.CreateAuthor(&author)
			if err != nil {
				return nil, err
			}
			authorID = newAuthor.ID
		} else {
			authorID = (*existingAuthor)[0].ID
		}

		// 同じauthorが複数回出てきたらスキップ
		if slices.Contains(authorIDs, authorID) {
			continue
		}
		bookAuthorRelations = append(
			bookAuthorRelations,
			entities.BookAuthor{
				BookID:   book.ID,
				AuthorID: authorID,
			},
		)
		authorIDs = append(authorIDs, authorID)
	}

	return &bookAuthorRelations, nil
}
