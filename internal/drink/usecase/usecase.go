package usecase

import (
	"goexample/internal/drink"
	"goexample/internal/drink/repository"
	"io"
)

type DrinkUsecase interface {
	CreateDrink(name, ingredients, description string, isAlcoholic bool, rating int, image io.ReadSeeker, imageName string) (*drink.Drink, error)
	GetAllDrinks() ([]drink.Drink, error)
	DeleteDrinkByID(id uint) error
}

type drinkUsecase struct {
	repo    repository.DrinkRepository
	storage drink.Storage
}

func NewDrinkUsecase(repo repository.DrinkRepository, storage drink.Storage) DrinkUsecase {
	return &drinkUsecase{repo: repo, storage: storage}
}

func (u *drinkUsecase) CreateDrink(name, ingredients, description string, isAlcoholic bool, rating int, image io.ReadSeeker, imageName string) (*drink.Drink, error) {

	contentType := "image/jpeg" // Ajuste conforme o tipo de imagem
	key := "images/drinks/" + imageName

	imageURL, err := u.storage.UploadFile(key, image, contentType)
	if err != nil {
		return nil, err
	}

	newDrink := &drink.Drink{
		Name:        name,
		Ingredients: ingredients,
		Description: description,
		IsAlcoholic: isAlcoholic,
		Rating:      rating,
		ImageURL:    imageURL,
	}
	err = u.repo.Create(newDrink)
	if err != nil {
		return nil, err
	}
	return newDrink, err
}

func (u *drinkUsecase) GetAllDrinks() ([]drink.Drink, error) {
	return u.repo.GetAll()
}

func (u *drinkUsecase) DeleteDrinkByID(id uint) error {
	return u.repo.DeleteByID(id)
}
