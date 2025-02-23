package services

import (
	"github.com/ttttai/golang/domain/entities"
	"github.com/ttttai/golang/domain/repositories"
)

type ISubjectService interface {
	CreateSubjects(subjects *[]entities.Subject) (*[]entities.Subject, error)
}

type SubjectService struct {
	subjectRepository repositories.ISubjectRepository
}

func NewSubjectService(subjectRepository repositories.ISubjectRepository) ISubjectService {
	return &SubjectService{
		subjectRepository: subjectRepository,
	}
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
