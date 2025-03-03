package domain

import (
	"strings"

	"github.com/google/uuid"
)

type User struct {
	id        uuid.UUID
	firstName string
	lastName  string
	email     string
	password  string
	salt      string
	role      Role
}

func NewUser(id uuid.UUID, firstName string, lastName string, email string, password string, salt string, role Role) (*User, error) {
	// Todo validate params

	return &User{
		id,
		firstName,
		lastName,
		email,
		password,
		salt,
		role,
	}, nil
}

func (u User) Id() uuid.UUID {
	return u.id
}

func (u User) FirstName() string {
	return u.firstName
}

func (u User) LastName() string {
	return u.lastName
}

func (u User) Email() string {
	return u.email
}

func (u User) Password() string {
	return u.password
}

func (u User) Salt() string {
	return u.salt
}

func (u User) Role() Role {
	return u.role
}

type Role int

const (
	RoleUser Role = iota
	RoleAdmin
)

func (r Role) String() string {
	return [2]string{"user", "admin"}[r]
}

func GetRole(s string) Role {
	switch strings.TrimSpace(strings.ToLower(s)) {
	case "admin":
		return RoleAdmin
	default:
		return RoleUser
	}
}
