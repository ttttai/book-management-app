package repositories

import (
	"github.com/ttttai/golang/domain/entities"
	"github.com/ttttai/golang/domain/repositories"
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
	var user entities.User

	result := ur.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (ur *UserRepository) Create(user *entities.User) (*entities.User, error) {
	result := ur.db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (ur *UserRepository) Update(user *entities.User) (*entities.User, error) {
	result := ur.db.Save(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (ur *UserRepository) Delete(id string) error {
	result := ur.db.Delete(&entities.User{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (ur *UserRepository) GetByName(name string) (*[]entities.User, error) {
	var users []entities.User

	result := ur.db.Where("name LIKE ?", "%"+name+"%").Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return &users, nil
}
