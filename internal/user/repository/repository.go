package repository

import "goexample/internal/user"

type UserRepository interface {
	Create(user *user.User) error
	GetAll() ([]user.User, error)
}
