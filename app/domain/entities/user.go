package entities

type User struct {
	id    string
	name  string
	email string
}

func NewUser(id, name string, email string) *User {
	return &User{
		id:    id,
		name:  name,
		email: email,
	}
}

func (u *User) GetID() string {
	return u.id
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) SetName(name string) {
	u.name = name
}
