package sqlstore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Seemann-ng/go-RESTapi/internal/app/model"
	"github.com/Seemann-ng/go-RESTapi/internal/app/store"
	"github.com/Seemann-ng/go-RESTapi/internal/app/store/sqlstore"
)

func TestUserRepository_Create(test *testing.T) {
	db, teardown := sqlstore.TestDB(test, databaseURL)
	defer teardown("users")

	storage := sqlstore.New(db)
	testUser := model.TestUser(test)

	assert.NoError(test, storage.User().Create(testUser))
	assert.NotNil(test, testUser)
}

func TestUserRepository_FindByEmail(test *testing.T) {
	db, teardown := sqlstore.TestDB(test, databaseURL)
	defer teardown("users")

	storage := sqlstore.New(db)
	email := "user@example.com"
	_, err := storage.User().FindByEmail(email)

	assert.EqualError(test, err, store.ErrRecordNotFound.Error())

	testUser := model.TestUser(test)
	testUser.Email = email
	if err = storage.User().Create(testUser); err != nil {
		test.Fatal(err)
	}

	testUser, err = storage.User().FindByEmail(email)

	assert.NoError(test, err)
	assert.NotNil(test, testUser)
}
