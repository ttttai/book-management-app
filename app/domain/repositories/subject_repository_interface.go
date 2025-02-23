package repositories

import "github.com/ttttai/golang/domain/entities"

type ISubjectRepository interface {
	CreateSubjects(subjects *[]entities.Subject) (*[]entities.Subject, error)
}
