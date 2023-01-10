package pg

import (
	"test-connect-db/server/repository"
	"test-connect-db/server/repository/users"

	"gorm.io/gorm"
)

// NewRepo new pg repo implimentation
func NewRepo(db *gorm.DB) *repository.Repo {
	return &repository.Repo{
		Users: users.NewPG(db),
	}
}
