package repository

import (
	"context"
	"database/sql"

	"github.com/ajaypanthagani/sizzlie/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// UserRepository is the interface for interacting with the users table.
type UserRepository interface {
	Create(user *models.User) (*models.User, error)
	Get(email string) (*models.User, error)
}

// UserRepositoryImpl is an implementation of the UserRepository interface.
type UserRepositoryImpl struct {
	db *sql.DB // Database connection
}

// NewUserRepository creates and returns a new UserRepositoryImpl instance.
func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}

// Create inserts a new user into the database.
func (repo *UserRepositoryImpl) Create(user *models.User) (*models.User, error) {
	err := user.Insert(context.Background(), repo.db, boil.Infer())
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Create inserts a new user into the database.
func (repo *UserRepositoryImpl) Get(email string) (*models.User, error) {
	user, err := models.Users(qm.Where("email = %s", email)).One(context.Background(), repo.db)
	if err != nil {
		return nil, err
	}
	return user, nil
}
