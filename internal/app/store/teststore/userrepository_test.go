package teststore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Seemann-ng/go-RESTapi/internal/app/model"
	"github.com/Seemann-ng/go-RESTapi/internal/app/store"
	"github.com/Seemann-ng/go-RESTapi/internal/app/store/teststore"
)

func TestUserRepository_Create(test *testing.T) {
	testStorage := teststore.New()
	user := model.TestUser(test)

	assert.NoError(test, testStorage.User().Create(user))
	assert.NotNil(test, user)
}

func TestUserRepository_FindByEmail(test *testing.T) {
	testStorage := teststore.New()
	email := "user@example.com"
	_, err := testStorage.User().FindByEmail(email)

	assert.EqualError(test, err, store.ErrRecordNotFound.Error())

	testUser := model.TestUser(test)
	testUser.Email = email
	if err = testStorage.User().Create(testUser); err != nil {
		test.Fatal(err)
	}

	testUser, err = testStorage.User().FindByEmail(email)

	assert.NoError(test, err)
	assert.NotNil(test, testUser)
}
