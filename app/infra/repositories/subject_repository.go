package repositories

import (
	"errors"

	"github.com/ttttai/golang/domain/entities"
	"github.com/ttttai/golang/domain/repositories"
	"gorm.io/gorm"
)

type SubjectRepository struct {
	db *gorm.DB
}

func NewSubjectRepository(db *gorm.DB) repositories.ISubjectRepository {
	return &SubjectRepository{
		db: db,
	}
}

func (sr *SubjectRepository) CreateSubject(subject *entities.Subject) (*entities.Subject, error) {
	result := sr.db.Create(subject)
	if result.Error != nil {
		return nil, result.Error
	}

	return subject, nil
}

func (sr *SubjectRepository) CreateSubjects(subjects *[]entities.Subject) (*[]entities.Subject, error) {
	result := sr.db.Create(subjects)
	if result.Error != nil {
		return nil, result.Error
	}

	return subjects, nil
}

func (sr *SubjectRepository) GetSubjectByName(name string) (*entities.Subject, error) {
	var subject entities.Subject

	result := sr.db.Where("subject_name = ?", name).First(&subject)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &subject, nil
}
