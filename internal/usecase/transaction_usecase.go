package usecase

import (
	"context"
	"errors"
	"math"
	"virtual-bank/internal/constants"
	"virtual-bank/internal/entity"
	"virtual-bank/internal/helpers"
	"virtual-bank/internal/models"
	"virtual-bank/internal/repository"

	"gorm.io/gorm"
)

type TransactionUsecase struct {
	db              *gorm.DB
	transactionRepo *repository.TransactionRepository
	accountRepo     *repository.AccountRepository
	log             *helpers.Logger
}

func NewTransactionUsecase(db *gorm.DB, transactionRepo *repository.TransactionRepository, accountRepo *repository.AccountRepository, log *helpers.Logger) *TransactionUsecase {
	return &TransactionUsecase{
		db:              db,
		transactionRepo: transactionRepo,
		log:             log,
		accountRepo:     accountRepo,
	}
}

// Return Code
//   - AccountNotFound
//   - DatabaseError
//   - NotEnoughBalance
//   - DestinationAccountNotFound
//   - SuccessInsert
func (c *TransactionUsecase) Create(ctx context.Context, dataCreate models.CreateTransaction) (*entity.Transaction, int) {
	tx := c.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	var account entity.Account
	if err := c.accountRepo.FindById(tx, &account, dataCreate.AccountId); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, constants.AccountNotFound
		}
		c.log.Error("%v", err)
		return nil, constants.DatabaseError
	}

	if account.Balance+dataCreate.Amount < 0 {
		return nil, constants.NotEnoughBalance
	}

	var destAccount entity.Account
	if err := c.accountRepo.FindById(tx, &destAccount, dataCreate.DestAccountId); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, constants.DestinationAccountNotFound
		}
		c.log.Error("%v", err)
		return nil, constants.DatabaseError
	}

	newTransaction := entity.Transaction{
		AccountId:       dataCreate.AccountId,
		TransactionType: dataCreate.TransactionType,
		DestAccountId:   dataCreate.DestAccountId,
		Description:     dataCreate.Description,
		Amount:          dataCreate.Amount,
		CurrentBalance:  account.Balance,
		Currency:        dataCreate.Currency,
		FinalBalance:    account.Balance + dataCreate.Amount,
	}

	account.Balance = account.Balance + dataCreate.Amount
	if err := c.accountRepo.Update(tx, &account); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, constants.AccountNotFound
		}
		c.log.Error("%v", err)
		return nil, constants.DatabaseError
	}

	destAccount.Balance = destAccount.Balance + int(math.Abs(float64(dataCreate.Amount)))
	if err := c.accountRepo.Update(tx, &destAccount); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, constants.DestinationAccountNotFound
		}
		c.log.Error("%v", err)
		return nil, constants.DatabaseError
	}

	if err := c.transactionRepo.Create(tx, &newTransaction); err != nil {
		c.log.Error("%v", err)
		return nil, constants.DatabaseError
	}

	if err := tx.Commit().Error; err != nil {
		c.log.Error("%v", err)
		return nil, constants.DatabaseError
	}
	return &newTransaction, constants.SuccessInsert
}

// Return Code
//   - RecordNotFound
//   - DatabaseError
//   - SuccessRead
func (c *TransactionUsecase) GetOne(ctx context.Context, id string) (*entity.Transaction, int) {
	tx := c.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	var transaction entity.Transaction

	if err := c.transactionRepo.FindById(tx, &transaction, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, constants.RecordNotFound
		}
		c.log.Error("%v", err)
		return nil, constants.DatabaseError
	}

	if err := tx.Commit().Error; err != nil {
		c.log.Error("%v", err)
		return nil, constants.DatabaseError
	}
	return &transaction, constants.SuccessRead
}

// Return Code
//   - DatabaseError
//   - SuccessRead
func (c *TransactionUsecase) GetAllPagination(ctx context.Context, accountId string, p repository.Pagination) ([]entity.Transaction, repository.MetaPagination, int) {
	tx := c.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	transaction, meta, err := c.transactionRepo.PaginationAggregation(tx.Model(&entity.Transaction{}).Where("account_id = ?", accountId), p)
	if err != nil {
		return transaction, meta, constants.DatabaseError
	}

	if err := tx.Commit().Error; err != nil {
		c.log.Error("%v", err)
		return nil, meta, constants.DatabaseError
	}
	return transaction, meta, constants.SuccessRead
}

// Return Code
//   - DatabaseError
//   - SuccessRead
func (c *TransactionUsecase) GetAllPaginationSubmissions(ctx context.Context, accountId string, p repository.Pagination) ([]entity.Transaction, repository.MetaPagination, int) {
	tx := c.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	transaction, meta, err := c.transactionRepo.PaginationAggregation(tx.Model(&entity.Transaction{}).Where("dest_account_id = ?", accountId), p)
	if err != nil {
		return transaction, meta, constants.DatabaseError
	}

	if err := tx.Commit().Error; err != nil {
		c.log.Error("%v", err)
		return nil, meta, constants.DatabaseError
	}
	return transaction, meta, constants.SuccessRead
}
