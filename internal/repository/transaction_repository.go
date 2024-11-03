package repository

import "virtual-bank/internal/entity"

type TransactionRepository struct {
	Repository[entity.Transaction]
}

func NewTransactionRepository() *TransactionRepository {
	return &TransactionRepository{}
}
