package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Seemann-ng/go-RESTapi/internal/app/model"
)

func TestUser_Validate(test *testing.T) {
	testCases := []struct {
		name    string
		user    func() *model.User
		isValid bool
	}{
		{
			name: "valid",
			user: func() *model.User {
				return model.TestUser(test)
			},
			isValid: true,
		},
		{
			name: "no email",
			user: func() *model.User {
				user := model.TestUser(test)
				user.Email = ""

				return user
			},
			isValid: false,
		},
		{
			name: "invalid email",
			user: func() *model.User {
				user := model.TestUser(test)
				user.Email = "invalid"

				return user
			},
			isValid: false,
		},
		{
			name: "empty password",
			user: func() *model.User {
				user := model.TestUser(test)
				user.Password = ""

				return user
			},
			isValid: false,
		},
		{
			name: "short password",
			user: func() *model.User {
				user := model.TestUser(test)
				user.Password = "abc"

				return user
			},
			isValid: false,
		},
		{
			name: "no password with encrypted password",
			user: func() *model.User {
				user := model.TestUser(test)
				user.Password = ""
				user.EncryptedPassword = "encryptedPassword"

				return user
			},
			isValid: true,
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if testCase.isValid {
				assert.NoError(test, testCase.user().Validate())
			} else {
				assert.Error(test, testCase.user().Validate())
			}
		})
	}
}

func TestUser_BeforeCreate(test *testing.T) {
	testUser := model.TestUser(test)

	assert.NoError(test, testUser.BeforeCreate())
	assert.NotEmpty(test, testUser.EncryptedPassword)
}
