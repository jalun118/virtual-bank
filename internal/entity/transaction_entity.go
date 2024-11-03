package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	ID              string    `json:"id" gorm:"primaryKey;size:50;unique"`
	AccountId       string    `json:"account_id"`
	TransactionType string    `json:"transaction_type"`
	DestAccountId   string    `json:"dest_account_id"`
	Description     string    `json:"description"`
	Amount          int       `json:"amount"`
	CurrentBalance  int       `json:"current_balance"`
	Currency        string    `json:"Currency"`
	FinalBalance    int       `json:"final_balance"`
	TransactionDate time.Time `json:"transaction_date"`
}

func (e *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	uid, err := uuid.NewV6()
	if err != nil {
		return err
	}
	e.ID = uid.String()
	e.TransactionDate = time.Now()
	return
}
