package repositories

import "github.com/ttttai/golang/domain/entities"

type IAuthenticationRepository interface {
	GetByEmail(email string) (*entities.Authentication, error)
	Create(authentication *entities.Authentication) (*entities.Authentication, error)
}
