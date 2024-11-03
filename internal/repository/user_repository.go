package repository

import (
	"virtual-bank/internal/entity"

	"gorm.io/gorm"
)

type UserRepository struct {
	Repository[entity.User]
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) FindByEmail(db *gorm.DB, entity *entity.User, email string) error {
	return db.Where("email = ?", email).Take(entity).Error
}
