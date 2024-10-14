package usecase

import (
	"goexample/internal/user"
	"goexample/internal/user/repository"
)

type UserUsecase interface {
	CreateUser(name string) (*user.User, error) // Retornar o usuário criado e o erro, se houver
	GetAllUsers() ([]user.User, error)
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}

func (u *userUsecase) CreateUser(name string) (*user.User, error) {
	newUser := &user.User{Name: name}
	err := u.repo.Create(newUser)
	return newUser, err // Agora retornamos o usuário criado e o erro
}

func (u *userUsecase) GetAllUsers() ([]user.User, error) {
	return u.repo.GetAll()
}
