package repositories

import (
	"github.com/ttttai/golang/domain/entities"
	"github.com/ttttai/golang/domain/repositories"
	"github.com/ttttai/golang/infra/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repositories.IUserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) GetById(id string) (*entities.User, error) {
	var user *models.User

	result := ur.db.First(user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return user.ToDomainModel(), nil
}

func (ur *UserRepository) Create(user *models.User) (*entities.User, error) {
	result := ur.db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user.ToDomainModel(), nil
}
