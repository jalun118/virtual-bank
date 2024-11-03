package repository

import (
	"errors"
	"virtual-bank/internal/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AccountRepository struct {
	Repository[entity.Account]
}

func NewAccountRepository() *AccountRepository {
	return &AccountRepository{}
}

func (r *AccountRepository) FoundRecordByUserId(db *gorm.DB, user_id uint) (bool, error) {
	var count int64
	if err := db.Model(&entity.Account{}).Where("user_id = ?", user_id).Count(&count).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (r *AccountRepository) FindByUserId(db *gorm.DB, entity *entity.Account, userId uint) error {
	return db.Where("user_id = ?", userId).Take(entity).Error
}

func (r *AccountRepository) FindByUserIdPreloadAll(db *gorm.DB, entity *entity.Account, userId uint) error {
	return db.Where("user_id = ?", userId).Preload(clause.Associations).Take(entity).Error
}
