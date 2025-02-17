package entities

import "time"

type User struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
}

func NewUser(id, name string, email string, create_at time.Time) *User {
	return &User{
		ID:        id,
		Name:      name,
		Email:     email,
		CreatedAt: create_at,
	}
}

func (u *User) GetID() string {
	return u.ID
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) SetName(name string) {
	u.Name = name
}
