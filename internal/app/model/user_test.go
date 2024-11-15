package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Seemann-ng/go-RESTapi/internal/app/model"
)

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		user    func() *model.User
		isValid bool
	}{
		{
			name: "valid",
			user: func() *model.User {
				return model.TestUser(t)
			},
			isValid: true,
		},
		{
			name: "no email",
			user: func() *model.User {
				user := model.TestUser(t)
				user.Email = ""

				return user
			},
			isValid: false,
		},
		{
			name: "invalid email",
			user: func() *model.User {
				user := model.TestUser(t)
				user.Email = "invalid"

				return user
			},
			isValid: false,
		},
		{
			name: "empty password",
			user: func() *model.User {
				user := model.TestUser(t)
				user.Password = ""

				return user
			},
			isValid: false,
		},
		{
			name: "short password",
			user: func() *model.User {
				user := model.TestUser(t)
				user.Password = "abc"

				return user
			},
			isValid: false,
		},
		{
			name: "no password with encrypted password",
			user: func() *model.User {
				user := model.TestUser(t)
				user.Password = ""
				user.EncryptedPassword = "encryptedPassword"

				return user
			},
			isValid: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.user().Validate())
			} else {
				assert.Error(t, tc.user().Validate())
			}
		})
	}
}

func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser(t)

	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)
}
