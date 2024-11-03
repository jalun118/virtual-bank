package entity

import (
	"time"
	"virtual-bank/internal/configs"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Session struct {
	ID        string    `json:"id" gorm:"primaryKey;size:50"`
	UserId    uint      `json:"user_id"`
	ExpiredAt time.Time `json:"expired_at"`
	CreatedAt time.Time `json:"created_at"`
}

func (e *Session) BeforeCreate(tx *gorm.DB) (err error) {
	uid, err := uuid.NewV6()
	if err != nil {
		return err
	}
	e.ID = uid.String()
	e.ExpiredAt = time.Now().Add(configs.MAX_AGE_SESSION)
	e.CreatedAt = time.Now()
	return
}
