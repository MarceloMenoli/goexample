// internal/user/usecase/usecase.go
package usecase

import (
	"goexample/internal/user"
	"goexample/internal/user/repository"
)

type UserUsecase interface {
	CreateUser(name string) error
	GetAllUsers() ([]user.User, error)
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}

func (u *userUsecase) CreateUser(name string) error {
	newUser := &user.User{Name: name}
	return u.repo.Create(newUser)
}

func (u *userUsecase) GetAllUsers() ([]user.User, error) {
	return u.repo.GetAll()
}
