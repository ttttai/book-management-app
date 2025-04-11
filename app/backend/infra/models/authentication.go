package models

import (
	"time"

	"github.com/ttttai/golang/domain/entities"
)

type Authentication struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `gorm:"<-:create"`
	UpdatedAt time.Time
}

func ToAuthenticationDomainModel(auth *Authentication) *entities.Authentication {
	return &entities.Authentication{
		ID:       auth.ID,
		Email:    auth.Email,
		Password: auth.Password,
	}
}

func FromAuthenticationDomainModel(auth *entities.Authentication) *Authentication {
	return &Authentication{
		ID:       auth.ID,
		Email:    auth.Email,
		Password: auth.Password,
	}
}
