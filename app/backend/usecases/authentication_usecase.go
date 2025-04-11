package usecases

import (
	"github.com/ttttai/golang/domain/entities"
	"github.com/ttttai/golang/domain/repositories"
)

type IAuthenticationUsecase interface {
	GetByEmail(email string) (*entities.Authentication, error)
	Create(authentication *entities.Authentication) (*entities.Authentication, error)
}

type AuthenticationUsecase struct {
	authenticationRepository repositories.IAuthenticationRepository
}

func NewAuthenticationUsecase(authenticationRepository repositories.IAuthenticationRepository) IAuthenticationUsecase {
	return &AuthenticationUsecase{
		authenticationRepository: authenticationRepository,
	}
}

func (au *AuthenticationUsecase) GetByEmail(email string) (*entities.Authentication, error) {
	authentication, err := au.authenticationRepository.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	return authentication, nil
}

func (au *AuthenticationUsecase) Create(authentication *entities.Authentication) (*entities.Authentication, error) {
	createdAuthentication, err := au.authenticationRepository.Create(authentication)
	if err != nil {
		return nil, err
	}

	return createdAuthentication, nil
}
