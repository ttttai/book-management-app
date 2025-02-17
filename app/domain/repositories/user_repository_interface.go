package repositories

import (
	"github.com/ttttai/golang/domain/entities"
)

type IUserRepository interface {
	GetUser(id string) (*entities.User, error)
	// UpdateUser(user *entities.User) (*entities.User, error)
}
