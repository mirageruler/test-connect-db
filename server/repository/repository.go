package repository

import "test-connect-db/server/repository/users"

type Repo struct {
	Users users.Store
}
