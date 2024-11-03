package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"size:256"`
	Password  string    `json:"password" gorm:"size:256"`
	Email     string    `json:"email" gorm:"size:256;unique"`
	FullName  string    `json:"full_name" gorm:"size:256"`
	BirthDate time.Time `json:"birth_date"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (e *User) BeforeCreate(tx *gorm.DB) (err error) {
	e.CreatedAt = time.Now()
	e.UpdatedAt = time.Now()
	return
}

func (e *User) BeforeSave(tx *gorm.DB) (err error) {
	e.UpdatedAt = time.Now()
	return
}
