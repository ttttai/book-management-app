package services

import (
	"github.com/ttttai/golang/domain/entities"
	"github.com/ttttai/golang/domain/repositories"
)

type ISubjectService interface {
	CreateSubject(subject *entities.Subject) (*entities.Subject, error)
	CreateSubjects(subjects *[]entities.Subject) (*[]entities.Subject, error)
	GetSubjectsByName(name string) (*[]entities.Subject, error)
	GetBookSubjectRelations(book *entities.Book, subjects *[]entities.Subject) (*[]entities.BookSubject, error)
}

type SubjectService struct {
	subjectRepository repositories.ISubjectRepository
}

func NewSubjectService(subjectRepository repositories.ISubjectRepository) ISubjectService {
	return &SubjectService{
		subjectRepository: subjectRepository,
	}
}

func (ss *SubjectService) CreateSubject(subject *entities.Subject) (*entities.Subject, error) {
	result, err := ss.subjectRepository.CreateSubject(subject)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (ss *SubjectService) CreateSubjects(subjects *[]entities.Subject) (*[]entities.Subject, error) {
	if len(*subjects) == 0 {
		return nil, nil
	}

	result, err := ss.subjectRepository.CreateSubjects(subjects)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (ss *SubjectService) GetSubjectsByName(name string) (*[]entities.Subject, error) {
	result, err := ss.subjectRepository.GetSubjectsByName(name)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (ss *SubjectService) GetBookSubjectRelations(book *entities.Book, subjects *[]entities.Subject) (*[]entities.BookSubject, error) {
	var bookSubjectRelations []entities.BookSubject
	for _, subject := range *subjects {
		if subject.SubjectName == "" {
			continue
		}

		var subjectID int
		existingSubject, _ := ss.GetSubjectsByName(subject.SubjectName)
		if len(*existingSubject) == 0 {
			newSubject, err := ss.CreateSubject(&subject)
			if err != nil {
				return nil, err
			}
			subjectID = newSubject.ID
		} else {
			subjectID = (*existingSubject)[0].ID
		}

		bookSubjectRelations = append(
			bookSubjectRelations,
			entities.BookSubject{
				BookID:    book.ID,
				SubjectID: subjectID,
			},
		)
	}

	return &bookSubjectRelations, nil
}
