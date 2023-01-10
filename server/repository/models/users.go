package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Users struct {
	ID      uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name    string
	Age     int
	IsAdmin bool
}

// BeforeCreate creates
func (u *Users) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewV4()
	return nil
}
