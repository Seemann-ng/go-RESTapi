package model

import (
	"testing"
)

// TestUser ...
func TestUser(test *testing.T) *User {
	return &User{
		Email:    "user@example.org",
		Password: "password",
	}
}
