// internal/user/repository/postgres/repository.go
package postgres

import (
	"goexample/internal/user"

	"gorm.io/gorm"
)

type postgresUserRepository struct {
	db *gorm.DB
}

func NewPostgresUserRepository(db *gorm.DB) *postgresUserRepository {
	return &postgresUserRepository{db: db}
}

func (r *postgresUserRepository) Create(user *user.User) error {
	return r.db.Create(user).Error
}

func (r *postgresUserRepository) GetAll() ([]user.User, error) {
	var users []user.User
	err := r.db.Find(&users).Error
	return users, err
}
