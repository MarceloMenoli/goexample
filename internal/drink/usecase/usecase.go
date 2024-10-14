package usecase

import (
	"goexample/internal/drink"
	"goexample/internal/drink/repository"
)

type DrinkUsecase interface {
	CreateDrink(name, ingredients, description string, isAlcoholic bool, rating int) (*drink.Drink, error)
	GetAllDrinks() ([]drink.Drink, error)
	DeleteDrinkByID(id uint) error
}

type drinkUsecase struct {
	repo repository.DrinkRepository
}

func NewDrinkUsecase(repo repository.DrinkRepository) DrinkUsecase {
	return &drinkUsecase{repo: repo}
}

func (u *drinkUsecase) CreateDrink(name, ingredients, description string, isAlcoholic bool, rating int) (*drink.Drink, error) {
	newDrink := &drink.Drink{
		Name:        name,
		Ingredients: ingredients,
		Description: description,
		IsAlcoholic: isAlcoholic,
		Rating:      rating,
	}
	err := u.repo.Create(newDrink)
	return newDrink, err
}

func (u *drinkUsecase) GetAllDrinks() ([]drink.Drink, error) {
	return u.repo.GetAll()
}

func (u *drinkUsecase) DeleteDrinkByID(id uint) error {
	return u.repo.DeleteByID(id)
}
