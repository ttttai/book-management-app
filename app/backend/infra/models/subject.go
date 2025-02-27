package models

import (
	"time"

	"github.com/ttttai/golang/domain/entities"
)

type Subject struct {
	ID          int `gorm:"primaryKey"`
	SubjectName string
	SubjectKana string
	CreatedAt   time.Time `gorm:"<-:create"`
	UpdatedAt   time.Time

	Books []Book `gorm:"many2many:book_authors;"`
}

func ToSubjectDomainModel(subject *Subject) *entities.Subject {
	return &entities.Subject{
		ID:          subject.ID,
		SubjectName: subject.SubjectName,
		SubjectKana: subject.SubjectKana,
	}
}

func ToSubjectDomainModels(subjects *[]Subject) *[]entities.Subject {
	subjectEntities := make([]entities.Subject, len(*subjects))
	for i, subject := range *subjects {
		subjectEntities[i] = *ToSubjectDomainModel(&subject)
	}
	return &subjectEntities
}

func FromSubjectDomainModel(subject *entities.Subject) *Subject {
	return &Subject{
		ID:          subject.ID,
		SubjectName: subject.SubjectName,
		SubjectKana: subject.SubjectKana,
	}
}

func FromSubjectDomainModels(subjects *[]entities.Subject) *[]Subject {
	subjectModels := make([]Subject, len(*subjects))
	for i, subject := range *subjects {
		subjectModels[i] = *FromSubjectDomainModel(&subject)
	}
	return &subjectModels
}
