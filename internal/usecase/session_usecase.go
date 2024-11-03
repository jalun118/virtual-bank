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

type SessionUsecase struct {
	db          *gorm.DB
	sessionRepo *repository.SessionRepository
	userRepo    *repository.UserRepository
	log         *helpers.Logger
}

func NewSessionUsecase(db *gorm.DB, sessionRepo *repository.SessionRepository, useRepo *repository.UserRepository, log *helpers.Logger) *SessionUsecase {
	return &SessionUsecase{
		db:          db,
		sessionRepo: sessionRepo,
		userRepo:    useRepo,
		log:         log,
	}
}

// Return Code
//   - UserNotFound
//   - DatabaseError
//   - SuccessInsert
func (c *SessionUsecase) Create(ctx context.Context, payload models.PayloadSession) (*entity.Session, int) {
	tx := c.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	found, err := c.userRepo.FoundRecordById(tx, payload.UserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, constants.UserNotFound
		}
		c.log.Error("%v", err)
		return nil, constants.DatabaseError
	}
	if !found {
		return nil, constants.UserNotFound
	}

	newSession := entity.Session{
		UserId: payload.UserId,
	}

	if err := c.sessionRepo.Create(tx, &newSession); err != nil {
		c.log.Error("%v", err)
		return nil, constants.DatabaseError
	}

	if err := tx.Commit().Error; err != nil {
		c.log.Error("%v", err)
		return nil, constants.DatabaseError
	}

	return &newSession, constants.SuccessInsert
}

// Return Code
//   - RecordNotFound
//   - DatabaseError
//   - SuccessRead
func (c *SessionUsecase) FindByIdAndNotExpired(ctx context.Context, id string) (*entity.Session, int) {
	tx := c.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	var session entity.Session

	if err := c.sessionRepo.FindByIdAndNoExpired(tx, &session, id); err != nil {
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
	return &session, constants.SuccessRead
}

// Return Code
//   - RecordNotFound
//   - DatabaseError
//   - SuccessDelete
func (c *SessionUsecase) DeleteById(ctx context.Context, id string) int {
	tx := c.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.sessionRepo.DeleteById(tx, &entity.Session{}, id); err != nil {
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

// Return Code
//   - RecordNotFound
//   - DatabaseError
//   - SuccessDelete
func (c *SessionUsecase) DeleteByExpired(ctx context.Context) int {
	tx := c.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.sessionRepo.DeleteAllSessionExpired(tx); err != nil {
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
