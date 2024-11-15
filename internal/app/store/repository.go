package store

import (
	"github.com/Seemann-ng/go-RESTapi/internal/app/model"
)

// UserRepository ...
type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}
