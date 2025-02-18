package repositories

import (
	"github.com/ttttai/golang/domain/entities"
	"github.com/ttttai/golang/infra/models"
)

type IUserRepository interface {
	GetById(id string) (*entities.User, error)
	Create(user *models.User) (*entities.User, error)
	// UpdateUser(user *entities.User) (*entities.User, error)
}
