package repositories

import (
	"github.com/ttttai/golang/domain/entities"
)

type IUserRepository interface {
	GetById(id string) (*entities.User, error)
	Create(user *entities.User) (*entities.User, error)
	Update(user *entities.User) (*entities.User, error)
	Delete(id string) error
	GetByName(name string) (*[]entities.User, error)
}
