package repositories

import "github.com/ttttai/golang/domain/entities"

type ISubjectRepository interface {
	CreateSubject(subject *entities.Subject) (*entities.Subject, error)
	CreateSubjects(subjects *[]entities.Subject) (*[]entities.Subject, error)
	GetSubjectByName(name string) (*entities.Subject, error)
}
