package users

import (
	"test-connect-db/server/repository/models"

	"gorm.io/gorm"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{
		db: db,
	}
}

func (pg *pg) Save(user *models.Users) error {
	return pg.db.Save(user).Error
}

func (pg *pg) OneByID(id string) (*models.Users, error) {
	var user models.Users
	return &user, pg.db.Where("id = ?", id).First(&user).Error
}

func (pg *pg) ManyByName(name string) ([]models.Users, error) {
	users := make([]models.Users, 0)
	return users, pg.db.Where("name = ?", name).Find(&users).Error
}
