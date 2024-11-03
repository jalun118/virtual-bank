package models

import "time"

type CreateAccount struct {
	UserId        uint      `json:"user_id"`
	AccountTypeId uint      `json:"account_type_id" validate:"required,number,min=1"`
	Balance       int       `json:"balance" validate:"number,min=0"`
	Pin           string    `json:"pin" validate:"required,numstring,len=6"`
	Currency      string    `json:"Currency" validate:"required,alphanum"`
	MotherBirth   time.Time `json:"mother_birth" validate:"required"`
}

type TopUpAccount struct {
	AccountId string `json:"account_id" validate:"numstring,len=12"`
	Amount    int    `json:"amount" validate:"number,min=0"`
}

type MatchPinAccount struct {
	Pin string `json:"pin" validate:"required,numstring,len=6"`
}
