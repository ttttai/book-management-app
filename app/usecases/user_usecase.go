package usecases

import (
	"github.com/ttttai/golang/domain/entities"
	"github.com/ttttai/golang/domain/repositories"
	"github.com/ttttai/golang/infra/models"
)

type IUserUsecase interface {
	GetById(id string) (*entities.User, error)
	Create(user *models.User) (*entities.User, error)
}

type UserUsecase struct {
	userRepository repositories.IUserRepository
}

func NweUserUsecase(userRepository repositories.IUserRepository) IUserUsecase {
	return &UserUsecase{
		userRepository: userRepository,
	}
}

func (uu *UserUsecase) GetById(id string) (*entities.User, error) {
	user, err := uu.userRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uu *UserUsecase) Create(user *models.User) (*entities.User, error) {
	newUser, err := uu.userRepository.Create(user)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}
