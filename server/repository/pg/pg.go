package pg

import (
	"database/sql"
	"embed"
	"fmt"
	"time"

	"test-connect-db/configs"

	"github.com/pressly/goose/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"

	_ "github.com/lib/pq"
)

func NewPostgresStore(cfg *configs.Config) Store {
	ds := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DatabaseUser, cfg.DatabasePassword,
		cfg.DatabaseHost, cfg.DatabasePort, cfg.DatabaseName,
	)

	conn, err := sql.Open("postgres", ds)
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(postgres.New(
		postgres.Config{Conn: conn}),
		&gorm.Config{
			Logger: gormlogger.Default.LogMode(gormlogger.Silent),
		})
	if err != nil {
		panic(err)
	}

	goose.SetBaseFS(embedMigrations)
	genericDB, err := db.DB()
	if err != nil {
		return nil
	}
	if err := goose.Up(genericDB, "migrations"); err != nil {
		panic(err)
	}

	return &store{
		database:     db,
		shutdownFunc: conn.Close,
	}
}

//go:embed migrations/*.sql
var embedMigrations embed.FS

// FinallyFunc function to finish a transaction
type FinallyFunc = func(error) error

// FinallyFuncWithRollbackTo
type FinallyFuncWithRollbackTo = func(string, error) error

type Store interface {
	DB() *gorm.DB
	NewTransaction() (Store, FinallyFunc)
	NewTransactionWithRollbackTo() (Store, FinallyFuncWithRollbackTo)
	Shutdown() error
}

// store is implimentation of repository
type store struct {
	database     *gorm.DB
	shutdownFunc func() error
}

// Shutdown close database connection
func (s *store) Shutdown() error {
	if s.shutdownFunc != nil {
		return s.shutdownFunc()
	}
	return nil
}

// DB database connection
func (s *store) DB() *gorm.DB {
	return s.database
}

// NewTransaction for database connection
func (s *store) NewTransaction() (newRepo Store, finallyFn FinallyFunc) {
	newDB := s.database.Begin()

	finallyFn = func(err error) error {
		time.Sleep(1 * time.Second)
		if err != nil {
			nErr := newDB.Rollback().Error
			if nErr != nil {
				return nErr
			}
			return err
		}

		cErr := newDB.Commit().Error
		if cErr != nil {
			return cErr
		}
		return nil
	}

	return &store{database: newDB}, finallyFn
}

// NewTransactionWithRollbackTo
func (s *store) NewTransactionWithRollbackTo() (newRepo Store, finallyFn FinallyFuncWithRollbackTo) {
	newDB := s.database.Begin()

	finallyFn = func(to string, err error) error {
		time.Sleep(1 * time.Second)
		if err != nil {
			if to == "ROLLBACK_ALL" {
				nErr := newDB.Rollback().Error
				if nErr != nil {
					return nErr
				}
				return err
			}

			nErr := newDB.RollbackTo(to).Error
			if nErr != nil {
				return nErr
			}
		}

		cErr := newDB.Commit().Error
		if cErr != nil {
			return cErr
		}
		return err
	}

	return &store{database: newDB}, finallyFn
}
