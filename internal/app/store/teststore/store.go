package teststore

import (
	"github.com/Seemann-ng/go-RESTapi/internal/app/model"
	"github.com/Seemann-ng/go-RESTapi/internal/app/store"

	_ "github.com/lib/pq"
)

// Store ...
type Store struct {
	userRepository *UserRepository
}

// New ...
func New() *Store {
	return &Store{}
}

// User ...
// store.User().Create()
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[string]*model.User),
	}

	return s.userRepository
}
