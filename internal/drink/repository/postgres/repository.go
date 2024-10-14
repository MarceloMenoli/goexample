package postgres

import (
	"goexample/internal/drink"

	"gorm.io/gorm"
)

type postgresDrinkRepository struct {
	db *gorm.DB
}

func NewPostgresDrinkRepository(db *gorm.DB) *postgresDrinkRepository {
	return &postgresDrinkRepository{db: db}
}

func (r *postgresDrinkRepository) Create(drink *drink.Drink) error {
	return r.db.Create(drink).Error
}

func (r *postgresDrinkRepository) GetAll() ([]drink.Drink, error) {
	var drinks []drink.Drink
	err := r.db.Find(&drinks).Error
	return drinks, err
}

func (r *postgresDrinkRepository) DeleteByID(id uint) error {
	return r.db.Delete(&drink.Drink{}, id).Error
}
