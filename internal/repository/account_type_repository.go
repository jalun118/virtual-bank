package repository

import (
	"virtual-bank/internal/entity"

	"gorm.io/gorm"
)

type AccountTypeRepository struct {
	Repository[entity.AccountType]
}

func NewAccountTypeRepository() *AccountTypeRepository {
	return &AccountTypeRepository{}
}

func (r *AccountTypeRepository) FindByName(db *gorm.DB, entity *entity.AccountType, name string) error {
	return db.Where("name = ?", name).Take(entity).Error
}
