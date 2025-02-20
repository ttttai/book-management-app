package entities

import "time"

type User struct {
	ID        uint
	Name      string
	Email     string
	CreatedAt time.Time
}

func NewUser(name string, email string) *User {
	return &User{
		Name:  name,
		Email: email,
	}
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetEmail() string {
	return u.Email
}
