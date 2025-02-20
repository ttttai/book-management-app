package usecases

import (
	"github.com/ttttai/golang/domain/entities"
	"github.com/ttttai/golang/domain/repositories"
	"github.com/ttttai/golang/infra/models"
	"github.com/ttttai/golang/usecases/dto"
)

type IUserUsecase interface {
	GetById(id string) (*entities.User, error)
	Create(user dto.CreateUserRequestParam) (*entities.User, error)
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

func (uu *UserUsecase) Create(user dto.CreateUserRequestParam) (*entities.User, error) {
	newUserModel := models.User{
		Name:  user.Name,
		Email: user.Email,
	}
	newUser, err := uu.userRepository.Create(&newUserModel)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}
