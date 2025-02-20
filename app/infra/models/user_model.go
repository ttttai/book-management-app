package models

import (
	"time"

	"github.com/ttttai/golang/domain/entities"
)

type User struct {
	ID        uint `gorm:"primary_key"`
	Name      string
	Email     string
	CreatedAt time.Time
}

func (u *User) ToDomainModel() *entities.User {
	return &entities.User{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
	}
}
