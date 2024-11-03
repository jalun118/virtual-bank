package repository

import (
	"time"
	"virtual-bank/internal/entity"

	"gorm.io/gorm"
)

type SessionRepository struct {
	Repository[entity.Session]
}

func NewSessionRepository() *SessionRepository {
	return &SessionRepository{}
}

func (r *SessionRepository) FindByUserId(db *gorm.DB, entity *entity.Session, id string) error {
	return db.Where("user_id = ?", id).Take(entity).Error
}

func (r *SessionRepository) FindByIdAndNoExpired(db *gorm.DB, entity *entity.Session, id string) error {
	return db.Where("id = ? AND expired_at >= ?", id, time.Now()).Take(entity).Error
}

func (r *SessionRepository) DeleteAllSessionExpired(db *gorm.DB) error {
	return db.Unscoped().Where("expired_at <= ?", time.Now()).Delete(&entity.Session{}).Error
}

func (r *SessionRepository) FindAllByUserId(db *gorm.DB, entity []*entity.Session, id uint) error {
	return db.Where("user_id = ?", id).Find(entity).Error
}
