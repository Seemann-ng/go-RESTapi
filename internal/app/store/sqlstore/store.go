package sqlstore

import (
	"database/sql"

	"github.com/Seemann-ng/go-RESTapi/internal/app/store"

	_ "github.com/lib/pq"
)

// Store ...
type Store struct {
	db             *sql.DB
	userRepository *UserRepository
}

// New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// User ...
// store.User().Create()
func (storage *Store) User() store.UserRepository {
	if storage.userRepository != nil {
		return storage.userRepository
	}

	storage.userRepository = &UserRepository{
		store: storage,
	}

	return storage.userRepository
}
