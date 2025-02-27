package models

import "github.com/ttttai/golang/domain/entities"

type BookSubject struct {
	BookID    int
	SubjectID int

	Book    Book    `gorm:"foreignKey:BookID;references:ID;constraint:OnDelete:CASCADE"`
	Subject Subject `gorm:"foreignKey:SubjectID;references:ID;constraint:OnDelete:CASCADE"`
}

func ToBookSubjectDomainModel(bookSubject *BookSubject) *entities.BookSubject {
	return &entities.BookSubject{
		BookID:    bookSubject.BookID,
		SubjectID: bookSubject.SubjectID,
	}
}

func ToBookSubjectDomainModels(bookSubjects *[]BookSubject) *[]entities.BookSubject {
	bookSubjectEntities := make([]entities.BookSubject, len(*bookSubjects))
	for i, bookSubject := range *bookSubjects {
		bookSubjectEntities[i] = *ToBookSubjectDomainModel(&bookSubject)
	}
	return &bookSubjectEntities
}

func FromBookSubjectDomainModel(bookSubject *entities.BookSubject) *BookSubject {
	return &BookSubject{
		BookID:    bookSubject.BookID,
		SubjectID: bookSubject.SubjectID,
	}
}

func FromBookSubjectDomainModels(bookSubjects *[]entities.BookSubject) *[]BookSubject {
	bookSubjectModels := make([]BookSubject, len(*bookSubjects))
	for i, bookSubject := range *bookSubjects {
		bookSubjectModels[i] = *FromBookSubjectDomainModel(&bookSubject)
	}
	return &bookSubjectModels
}
