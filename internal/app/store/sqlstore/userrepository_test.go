package sqlstore_test

import (
	"github.com/GlebMoskalev/http-rest-api/internal/app/model"
	"github.com/GlebMoskalev/http-rest-api/internal/app/store"
	"github.com/GlebMoskalev/http-rest-api/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	db, tearDown := sqlstore.TestDb(t, databaseURL)
	defer tearDown("users")

	s := sqlstore.New(db)
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, tearDown := sqlstore.TestDb(t, databaseURL)
	defer tearDown("users")

	s := sqlstore.New(db)

	email := "testuser@example.com"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email
	s.User().Create(u)

	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_Find(t *testing.T) {
	db, tearDown := sqlstore.TestDb(t, databaseURL)
	defer tearDown("users")

	s := sqlstore.New(db)
	u1 := model.TestUser(t)
	s.User().Create(u1)
	u2, err := s.User().Find(u1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}
