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
	var user models.User

	result := ur.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return models.ToUserDomainModel(&user), nil
}

func (ur *UserRepository) Create(user *entities.User) (*entities.User, error) {
	userModel := models.FromUserDomainModel(user)
	result := ur.db.Create(userModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return models.ToUserDomainModel(userModel), nil
}

func (ur *UserRepository) Update(user *entities.User) (*entities.User, error) {
	userModel := models.FromUserDomainModel(user)
	result := ur.db.Save(userModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return models.ToUserDomainModel(userModel), nil
}

func (ur *UserRepository) Delete(id string) error {
	result := ur.db.Delete(&entities.User{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (ur *UserRepository) GetByName(name string) (*[]entities.User, error) {
	var users []models.User

	result := ur.db.Where("name LIKE ?", "%"+name+"%").Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return models.ToUserDomainModels(&users), nil
}
