package repository

import "goexample/internal/drink"

type DrinkRepository interface {
	Create(drink *drink.Drink) error
	GetAll() ([]drink.Drink, error)
	DeleteByID(id uint) error
}
