package models

type CreateTransaction struct {
	AccountId       string `json:"account_id"`
	DestAccountId   string `json:"dest_account_id" validate:"required,alphanumspecial"`
	TransactionType string `json:"transaction_type" validate:"required,alphanumspecial"`
	Amount          int    `json:"amount" validate:"required,number"`
	Currency        string `json:"Currency" validate:"required,alphanumspecial"`
	Description     string `json:"description" validate:"required"`
}
