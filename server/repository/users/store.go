package users

import "test-connect-db/server/repository/models"

type Store interface {
	Save(*models.Users) error
	OneByID(string) (*models.Users, error)
	ManyByName(string) ([]models.Users, error)
}
