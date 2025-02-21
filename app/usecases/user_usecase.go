package usecases

import (
	"github.com/ttttai/golang/domain/entities"
	"github.com/ttttai/golang/domain/repositories"
	"github.com/ttttai/golang/usecases/dto"
)

type IUserUsecase interface {
	GetById(id string) (*entities.User, error)
	Create(user dto.CreateUserRequestParam) (*entities.User, error)
	Update(id dto.UpdateUserRequestPathParam, data dto.UpdateUserRequestBodyParam) (*entities.User, error)
	Delete(id string) error
	GetByName(name string) (*[]entities.User, error)
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
	newUser := entities.NewUser(user.Name, user.Email)
	result, err := uu.userRepository.Create(newUser)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (uu *UserUsecase) Update(id dto.UpdateUserRequestPathParam, data dto.UpdateUserRequestBodyParam) (*entities.User, error) {
	user, err := uu.userRepository.GetById(id.ID)
	if err != nil {
		return nil, err
	}

	user.Name = data.Name
	user.Email = data.Email
	result, err := uu.userRepository.Update(user)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (uu *UserUsecase) Delete(id string) error {
	err := uu.userRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (uu *UserUsecase) GetByName(name string) (*[]entities.User, error) {
	users, err := uu.userRepository.GetByName(name)
	if err != nil {
		return nil, err
	}

	return users, nil
}
