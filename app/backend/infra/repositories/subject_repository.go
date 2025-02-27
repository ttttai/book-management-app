package repositories

import (
	"errors"

	"github.com/ttttai/golang/domain/entities"
	"github.com/ttttai/golang/domain/repositories"
	"github.com/ttttai/golang/infra/models"
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
	subjectModel := models.FromSubjectDomainModel(subject)
	result := sr.db.Create(subjectModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return models.ToSubjectDomainModel(subjectModel), nil
}

func (sr *SubjectRepository) CreateSubjects(subjects *[]entities.Subject) (*[]entities.Subject, error) {
	subjectModels := models.FromSubjectDomainModels(subjects)
	result := sr.db.Create(subjectModels)
	if result.Error != nil {
		return nil, result.Error
	}

	return models.ToSubjectDomainModels(subjectModels), nil
}

func (sr *SubjectRepository) GetSubjectsByName(name string) (*[]entities.Subject, error) {
	var subject []models.Subject

	result := sr.db.Where("subject_name = ?", name).Find(&subject)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return models.ToSubjectDomainModels(&subject), nil
}
