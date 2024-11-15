package sqlstore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Seemann-ng/go-RESTapi/internal/app/model"
	"github.com/Seemann-ng/go-RESTapi/internal/app/store"
	"github.com/Seemann-ng/go-RESTapi/internal/app/store/sqlstore"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	email := "user@example.com"
	_, err := s.User().FindByEmail(email)

	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email
	if err = s.User().Create(u); err != nil {
		t.Fatal(err)
	}

	u, err = s.User().FindByEmail(email)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}
