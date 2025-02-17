package models

import (
	"time"

	"github.com/ttttai/golang/domain/entities"
)

type User struct {
	ID        string    `json:"id" binding:"required"`
	Name      string    `json:"name" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *User) ToDomainEntity() *entities.User {
	return entities.NewUser(u.ID, u.Name, u.Email)
}
