package models

import (
	"time"

	"github.com/ttttai/golang/domain/entities"
)

type User struct {
	ID        int `gorm:"primaryKey"`
	Name      string
	Email     string
	CreatedAt time.Time `gorm:"<-:create"`
	UpdatedAt time.Time
}

func ToUserDomainModel(user *User) *entities.User {
	return &entities.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

func ToUserDomainModels(users *[]User) *[]entities.User {
	userEntities := make([]entities.User, len(*users))
	for i, user := range *users {
		userEntities[i] = *ToUserDomainModel(&user)
	}
	return &userEntities
}

func FromUserDomainModel(user *entities.User) *User {
	return &User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

func FromUserDomainModels(users *[]entities.User) *[]User {
	userModels := make([]User, len(*users))
	for i, user := range *users {
		userModels[i] = *FromUserDomainModel(&user)
	}
	return &userModels
}
