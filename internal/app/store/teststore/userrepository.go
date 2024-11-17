package teststore

import (
	"github.com/Seemann-ng/go-RESTapi/internal/app/model"
	"github.com/Seemann-ng/go-RESTapi/internal/app/store"
)

// UserRepository ...
type UserRepository struct {
	store *Store
	users map[string]*model.User
}

// Create ...
func (repo *UserRepository) Create(user *model.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	if err := user.BeforeCreate(); err != nil {
		return err
	}

	repo.users[user.Email] = user
	user.ID = len(repo.users)

	return nil
}

// FindByEmail ...
func (repo *UserRepository) FindByEmail(email string) (*model.User, error) {
	user, ok := repo.users[email]
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return user, nil
}
