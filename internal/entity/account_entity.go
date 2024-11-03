package entity

import (
	"errors"
	"time"
	"virtual-bank/internal/configs"
	"virtual-bank/internal/helpers"

	"gorm.io/gorm"
)

type Account struct {
	ID            string      `json:"id" gorm:"primaryKey;unique;size:12"`
	UserId        uint        `json:"user_id"`
	User          User        `json:"user" gorm:"foreignKey:UserId;references:ID"`
	AccountTypeId uint        `json:"account_type_id"`
	AccountType   AccountType `json:"account_type" gorm:"foreignKey:AccountTypeId;references:ID"`
	Currency      string      `json:"Currency"`
	Balance       int         `json:"balance"`
	Pin           string      `json:"pin"`
	MotherBirth   time.Time   `json:"mother_birth"`
	CreatedAt     time.Time   `json:"created_at"`
}

func (e *Account) BeforeCreate(tx *gorm.DB) (err error) {
	var count int64
	countLoop := 0

	for {
		e.ID = configs.PROVIDER_NUMBER + helpers.Generator.GenerateNumberString(9)

		find := tx.Where(&e, "id").Count(&count)

		errDb, rowAffected := find.Error, find.RowsAffected

		if errors.Is(errDb, gorm.ErrRecordNotFound) || rowAffected == 0 || count == 0 {
			break
		}

		if countLoop > 10 {
			return errors.New("infinite loop")
		}

		countLoop++
	}

	return
}
