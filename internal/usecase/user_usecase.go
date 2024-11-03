package usecase

import (
	"context"
	"errors"
	"virtual-bank/internal/constants"
	"virtual-bank/internal/entity"
	"virtual-bank/internal/helpers"
	"virtual-bank/internal/models"
	"virtual-bank/internal/repository"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserUsecase struct {
	db       *gorm.DB
	userRepo *repository.UserRepository
	log      *helpers.Logger
}

func NewUserUsecase(db *gorm.DB, userRepo *repository.UserRepository, log *helpers.Logger) *UserUsecase {
	return &UserUsecase{
		db:       db,
		log:      log,
		userRepo: userRepo,
	}
}

// Return Code
//   - DuplicateRecord
//   - DatabaseError
//   - InternalServerError
//   - SuccessInsert
func (c *UserUsecase) Create(ctx context.Context, dataUser models.CreateUser) (*entity.User, int) {
	tx := c.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	errFind := c.userRepo.FindByEmail(tx, &entity.User{}, dataUser.Email)
	if errFind == nil {
		return nil, constants.DuplicateRecord
	}
	if !errors.Is(errFind, gorm.ErrRecordNotFound) {
		c.log.Error("%v", errFind)
		return nil, constants.DatabaseError
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(dataUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.log.Error("%v", errFind)
		return nil, constants.InternalServerError
	}

	newUser := entity.User{
		Username:  dataUser.Username,
		Password:  string(hashPassword),
		Email:     dataUser.Email,
		FullName:  dataUser.FullName,
		BirthDate: dataUser.BirthDate,
	}

	if err := c.userRepo.Create(tx, &newUser); err != nil {
		c.log.Error("%v", err)
		return nil, constants.DatabaseError
	}

	if err := tx.Commit().Error; err != nil {
		c.log.Error("%v", err)
		return nil, constants.DatabaseError
	}

	return &newUser, constants.SuccessInsert
}

// Return Code
//   - RecordNotFound
//   - DatabaseError
//   - SuccessUpdate
func (c *UserUsecase) Update(ctx context.Context, id uint, dataUpdate models.UpdateUser) (*entity.User, int) {
	tx := c.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	var user entity.User

	if err := c.userRepo.FindById(tx, &user, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, constants.RecordNotFound
		}

		c.log.Error("%v", err)
		return nil, constants.DatabaseError
	}

	if dataUpdate.FullName != "" {
		user.FullName = dataUpdate.FullName
	}
	if dataUpdate.Username != "" {
		user.Username = dataUpdate.Username
	}
	if !dataUpdate.BirthDate.IsZero() {
		user.BirthDate = dataUpdate.BirthDate
	}

	if err := c.userRepo.Update(tx, &user); err != nil {
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

	return &user, constants.SuccessUpdate
}

// Return Code
//   - RecordNotFound
//   - DatabaseError
//   - PasswordNotMatch
//   - SuccessRead
func (c *UserUsecase) PasswordMatch(ctx context.Context, email string, password string) (bool, int) {
	tx := c.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	var user entity.User

	if err := c.userRepo.FindByEmail(tx, &user, email); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, constants.RecordNotFound
		}

		c.log.Error("%v", err)
		return false, constants.DatabaseError
	}

	isMatch := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if isMatch != nil {
		return false, constants.PasswordNotMatch
	}

	if err := tx.Commit().Error; err != nil {
		c.log.Error("%v", err)
		return true, constants.DatabaseError
	}

	return true, constants.SuccessRead
}

// Return Code
//   - RecordNotFound
//   - DatabaseError
//   - SuccessRead
func (c *UserUsecase) GetByEmail(ctx context.Context, email string) (*entity.User, int) {
	tx := c.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	var user entity.User

	if err := c.userRepo.FindByEmail(tx, &user, email); err != nil {
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

	return &user, constants.SuccessRead
}

// Return Code
//   - RecordNotFound
//   - DatabaseError
//   - SuccessRead
func (c *UserUsecase) GetById(ctx context.Context, id uint) (*entity.User, int) {
	tx := c.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	var user entity.User

	if err := c.userRepo.FindById(tx, &user, id); err != nil {
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

	return &user, constants.SuccessRead
}

// Return Code
//   - RecordNotFound
//   - DatabaseError
//   - SuccessUpdate
func (c *UserUsecase) SetStatusUser(ctx context.Context, id uint, status bool) (*entity.User, int) {
	tx := c.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	var user entity.User

	if err := c.userRepo.FindById(tx, &user, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, constants.RecordNotFound
		}

		c.log.Error("%v", err)
		return nil, constants.DatabaseError
	}

	if status {
		user.Status = 1
	}

	if !status {
		user.Status = 0
	}

	if err := c.userRepo.Update(tx, &user); err != nil {
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

	return &user, constants.SuccessUpdate
}

// Return Code
//   - RecordNotFound
//   - DatabaseError
//   - PasswordNotMatch
//   - InternalServerError
//   - SuccessUpdate
func (c *UserUsecase) ChangePassword(ctx context.Context, id uint, dataChange models.ChangePassword) (*entity.User, int) {
	tx := c.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	var user entity.User

	if err := c.userRepo.FindById(tx, &user, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, constants.RecordNotFound
		}

		c.log.Error("%v", err)
		return nil, constants.DatabaseError
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dataChange.OldPassword)); err != nil {
		return nil, constants.PasswordNotMatch
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(dataChange.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		c.log.Error("%v", err)
		return nil, constants.InternalServerError
	}

	user.Password = string(hash)

	if err := c.userRepo.Update(tx, &user); err != nil {
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

	return &user, constants.SuccessUpdate
}

// Return Code
//   - RecordNotFound
//   - DatabaseError
//   - SuccessDelete
func (c *UserUsecase) DeleteById(ctx context.Context, id uint) int {
	tx := c.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.userRepo.DeleteById(tx, &entity.User{}, id); err != nil {
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
