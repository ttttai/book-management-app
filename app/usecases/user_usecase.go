package usecases

import (
	"github.com/ttttai/golang/domain/entities"
	"github.com/ttttai/golang/domain/repositories"
)

type IUserUsecase interface {
	GetUser(id string) (*entities.User, error)
}

type UserUsecase struct {
	userRepository repositories.IUserRepository
}

func NweUserUsecase(userRepository repositories.IUserRepository) IUserUsecase {
	return &UserUsecase{
		userRepository: userRepository,
	}
}

func (uu *UserUsecase) GetUser(id string) (*entities.User, error) {
	user, err := uu.userRepository.GetUser(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
