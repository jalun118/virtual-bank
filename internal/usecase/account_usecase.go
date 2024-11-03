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

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AccountUsecase struct {
	db              *gorm.DB
	accountRepo     *repository.AccountRepository
	accountTypeRepo *repository.AccountTypeRepository
	userRepo        *repository.UserRepository
	log             *helpers.Logger
}

func NewAccountUsecase(db *gorm.DB, accountRepo *repository.AccountRepository, accountTypeRepo *repository.AccountTypeRepository, userRepo *repository.UserRepository, log *helpers.Logger) *AccountUsecase {
	return &AccountUsecase{
		db:              db,
		accountRepo:     accountRepo,
		accountTypeRepo: accountTypeRepo,
		userRepo:        userRepo,
		log:             log,
	}
}

// Return Code
//   - DatabaseError
//   - UserNotFound
//   - AccountTypeNotFound
//   - InternalServerError
//   - SuccessInsert
func (c *AccountUsecase) Create(ctx context.Context, dataCreate models.CreateAccount) (*entity.Account, int) {
	tx := c.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	foundUser, err := c.userRepo.FoundRecordById(tx, dataCreate.UserId)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			c.log.Error("%v", err)
			return nil, constants.DatabaseError
		}
	}
	if !foundUser {
		return nil, constants.UserNotFound
	}

	foundAccountType, err := c.accountTypeRepo.FoundRecordById(tx, dataCreate.AccountTypeId)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			c.log.Error("%v", err)
			return nil, constants.DatabaseError
		}
	}
	if !foundAccountType {
		return nil, constants.AccountTypeNotFound
	}

	foundAccount, err := c.accountRepo.FoundRecordByUserId(tx, dataCreate.UserId)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			c.log.Error("%v", err)
			return nil, constants.DatabaseError
		}
	}
	if foundAccount {
		return nil, constants.DuplicateRecord
	}

	hashPin, err := bcrypt.GenerateFromPassword([]byte(dataCreate.Pin), bcrypt.DefaultCost)
	if err != nil {
		c.log.Error("%v", err)
		return nil, constants.InternalServerError
	}

	newAccount := entity.Account{
		UserId:        dataCreate.UserId,
		AccountTypeId: dataCreate.AccountTypeId,
		Balance:       dataCreate.Balance,
		Pin:           string(hashPin),
		MotherBirth:   dataCreate.MotherBirth,
		Currency:      dataCreate.Currency,
	}

	if err := c.accountRepo.Create(tx, &newAccount); err != nil {
		c.log.Error("%v", err)
		return nil, constants.DatabaseError
	}

	if err := tx.Commit().Error; err != nil {
		c.log.Error("%v", err)
		return nil, constants.DatabaseError
	}

	return &newAccount, constants.SuccessInsert
}

// Return Code
//   - RecordNotFound
//   - DatabaseError
//   - SuccessRead
func (c *AccountUsecase) GetById(ctx context.Context, id string) (*entity.Account, int) {
	tx := c.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	var account entity.Account

	if err := c.accountRepo.FindByIdPreloadAll(tx, &account, id); err != nil {
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

	return &account, constants.SuccessRead
}

// Return Code
//   - RecordNotFound
//   - DatabaseError
//   - SuccessUpdate
func (c *AccountUsecase) TopUpAccount(ctx context.Context, dataTopUp models.TopUpAccount) (*entity.Account, int) {
	tx := c.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	var account entity.Account

	if err := c.accountRepo.FindById(tx, &account, dataTopUp.AccountId); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, constants.RecordNotFound
		}
		c.log.Error("%v", err)
		return nil, constants.DatabaseError
	}

	account.Balance = account.Balance + int(math.Abs(float64(dataTopUp.Amount)))

	if err := c.accountRepo.Update(tx, &account); err != nil {
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

	return &account, constants.SuccessUpdate
}

// Return Code
//   - RecordNotFound
//   - DatabaseError
//   - SuccessRead
func (c *AccountUsecase) GetByUserId(ctx context.Context, userId uint) (*entity.Account, int) {
	tx := c.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	var account entity.Account

	if err := c.accountRepo.FindByUserIdPreloadAll(tx, &account, userId); err != nil {
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

	return &account, constants.SuccessRead
}

// Return Code
//   - RecordNotFound
//   - DatabaseError
//   - PinNotMatch
//   - SuccessRead
func (c *AccountUsecase) MatchPin(ctx context.Context, id string, pin string) (bool, int) {
	tx := c.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	var account entity.Account

	if err := c.accountRepo.FindById(tx, &account, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, constants.RecordNotFound
		}
		c.log.Error("%v", err)
		return false, constants.DatabaseError
	}

	if err := bcrypt.CompareHashAndPassword([]byte(account.Pin), []byte(pin)); err != nil {
		return false, constants.PinNotMatch
	}

	if err := tx.Commit().Error; err != nil {
		c.log.Error("%v", err)
		return false, constants.DatabaseError
	}

	return true, constants.SuccessRead
}

// Return Code
//   - RecordNotFound
//   - DatabaseError
//   - PinNotMatch
//   - SuccessRead
func (c *AccountUsecase) MatchPinByUserId(ctx context.Context, userId uint, pin string) (bool, int) {
	tx := c.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	var account entity.Account

	if err := c.accountRepo.FindByUserId(tx, &account, userId); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, constants.RecordNotFound
		}
		c.log.Error("%v", err)
		return false, constants.DatabaseError
	}

	if err := bcrypt.CompareHashAndPassword([]byte(account.Pin), []byte(pin)); err != nil {
		return false, constants.PinNotMatch
	}

	if err := tx.Commit().Error; err != nil {
		c.log.Error("%v", err)
		return false, constants.DatabaseError
	}

	return true, constants.SuccessRead
}

// Return Code
//   - DatabaseError
//   - AccountTypeNotFound
//   - RecordNotFound
//   - SuccessUpdate
func (c *AccountUsecase) UpdateAccountType(ctx context.Context, id string, typeId uint) (*entity.Account, int) {
	tx := c.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	foundAccountType, err := c.accountTypeRepo.FoundRecordById(tx, typeId)
	if err != nil {
		c.log.Error("%v", err)
		return nil, constants.DatabaseError
	}
	if !foundAccountType {
		return nil, constants.AccountTypeNotFound
	}

	var account entity.Account

	if err := c.accountRepo.FindById(tx, &account, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, constants.RecordNotFound
		}
		c.log.Error("%v", err)
		return nil, constants.DatabaseError
	}

	account.AccountTypeId = typeId

	if err := c.accountRepo.Update(tx, &account); err != nil {
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

	return &account, constants.SuccessUpdate
}

// Return Code
//   - RecordNotFound
//   - DatabaseError
//   - SuccessDelete
func (c *AccountUsecase) DeleteById(ctx context.Context, id string) int {
	tx := c.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.accountRepo.DeleteById(tx, &entity.Account{}, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return constants.RecordNotFound
		}
		c.log.Error("%v", err)
		return constants.DatabaseError
	}

	if err := tx.Commit().Error; err != nil {
		c.log.Error("%v", err)
		return constants.DatabaseError
	}

	return constants.SuccessDelete
}
