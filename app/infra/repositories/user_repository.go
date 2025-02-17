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

func (ur *UserRepository) GetUser(id string) (*entities.User, error) {
	var user models.User

	result := ur.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return user.ToDomainEntity(), nil
}
