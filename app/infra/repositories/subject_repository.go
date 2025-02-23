package repositories

import (
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

func (sr *SubjectRepository) CreateSubjects(subjects *[]entities.Subject) (*[]entities.Subject, error) {
	result := sr.db.Create(subjects)
	if result.Error != nil {
		return nil, result.Error
	}

	return subjects, nil
}
