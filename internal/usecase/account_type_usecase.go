package usecase

import (
	"context"
	"errors"
	"virtual-bank/internal/constants"
	"virtual-bank/internal/entity"
	"virtual-bank/internal/helpers"
	"virtual-bank/internal/models"
	"virtual-bank/internal/repository"

	"gorm.io/gorm"
)

type AccountTypeUsecase struct {
	db              *gorm.DB
	accountTypeRepo *repository.AccountTypeRepository
	log             *helpers.Logger
}

func NewAccountTypeUsecase(db *gorm.DB, accountTypeRepo *repository.AccountTypeRepository, log *helpers.Logger) *AccountTypeUsecase {
	return &AccountTypeUsecase{
		db:              db,
		accountTypeRepo: accountTypeRepo,
		log:             log,
	}
}

// Return Code
//   - DatabaseError
//   - DuplicateRecord
//   - SuccessInsert
func (c *AccountTypeUsecase) CreateType(ctx context.Context, dataCreate models.CreateAccountType) (*entity.AccountType, int) {
	tx := c.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	var typeAccount entity.AccountType

	errFind := c.accountTypeRepo.FindByName(tx, &entity.AccountType{}, dataCreate.Name)
	if errFind != nil {
		if !errors.Is(errFind, gorm.ErrRecordNotFound) {
			return nil, constants.DatabaseError
		}
	}
	if typeAccount.Name != "" || errFind == nil {
		return nil, constants.DuplicateRecord
	}

	newType := entity.AccountType{
		Name: dataCreate.Name,
	}

	if err := c.accountTypeRepo.Create(tx, &newType); err != nil {
		c.log.Error("%v", err)
		return nil, constants.DatabaseError
	}

	if err := tx.Commit().Error; err != nil {
		c.log.Error("%v", err)
		return nil, constants.DatabaseError
	}

	return &newType, constants.SuccessInsert
}

// Return Code
//   - RecordNotFound
//   - DatabaseError
//   - SuccessRead
func (c *AccountTypeUsecase) GetById(ctx context.Context, id uint) (*entity.AccountType, int) {
	tx := c.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	var accountType entity.AccountType

	if err := c.accountTypeRepo.FindById(tx, &accountType, id); err != nil {
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

	return &accountType, constants.SuccessRead
}

// Return Code
//   - RecordNotFound
//   - DatabaseError
//   - SuccessRead
func (c *AccountTypeUsecase) GetAll(ctx context.Context) ([]entity.AccountType, int) {
	tx := c.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	datas, errFind := c.accountTypeRepo.GetAll(tx)
	if errFind != nil {
		if errors.Is(errFind, gorm.ErrRecordNotFound) {
			return nil, constants.RecordNotFound
		}
		c.log.Error("%v", errFind)
		return nil, constants.DatabaseError
	}

	if err := tx.Commit().Error; err != nil {
		c.log.Error("%v", err)
		return nil, constants.DatabaseError
	}

	return datas, constants.SuccessRead
}

// Return Code
//   - RecordNotFound
//   - DatabaseError
//   - SuccessUpdate
func (c *AccountTypeUsecase) Update(ctx context.Context, id uint, dataUpdate models.UpdateAccountType) (*entity.AccountType, int) {
	tx := c.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	var accountType entity.AccountType

	if err := c.accountTypeRepo.FindById(tx, &accountType, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, constants.RecordNotFound
		}
		c.log.Error("%v", err)
		return nil, constants.DatabaseError
	}

	accountType.Name = dataUpdate.Name

	if err := c.accountTypeRepo.Update(tx, &accountType); err != nil {
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

	return &accountType, constants.SuccessUpdate
}

// Return Code
//   - RecordNotFound
//   - DatabaseError
//   - SuccessDelete
func (c *AccountTypeUsecase) DeleteById(ctx context.Context, id uint) int {
	tx := c.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.accountTypeRepo.DeleteById(tx, &entity.AccountType{}, id); err != nil {
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
