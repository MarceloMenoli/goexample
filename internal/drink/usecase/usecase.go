package usecase

import (
	"goexample/internal/drink"
	"goexample/internal/drink/repository"
	"io"
	"os"
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

	contentType := "image/jpeg"
	key := "images/drinks/" + imageName

	err := u.storage.UploadFile(key, image, contentType)
	if err != nil {
		return nil, err
	}

	newDrink := &drink.Drink{
		Name:        name,
		Ingredients: ingredients,
		Description: description,
		IsAlcoholic: isAlcoholic,
		Rating:      rating,
		ImageURL:    key,
	}
	err = u.repo.Create(newDrink)
	if err != nil {
		return nil, err
	}
	return newDrink, err
}

func (u *drinkUsecase) GetAllDrinks() ([]drink.Drink, error) {
	drinks, err := u.repo.GetAll()
	if err != nil {
		return nil, err
	}

	baseURL := os.Getenv("IMAGE_BASE_URL")
	if baseURL == "" {
		baseURL = "https://pub-6b789ba0558c422f921e39eee0f551a6.r2.dev/"
	}

	for i, drink := range drinks {
		drinks[i].ImageURL = baseURL + drink.ImageURL
	}

	return drinks, nil
}

func (u *drinkUsecase) DeleteDrinkByID(id uint) error {
	return u.repo.DeleteByID(id)
}
