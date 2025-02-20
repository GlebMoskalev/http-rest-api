package store_test

import (
	"github.com/GlebMoskalev/http-rest-api/internal/app/model"
	"github.com/GlebMoskalev/http-rest-api/internal/app/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s, tearDown := store.TestStore(t, databaseURL)
	defer tearDown("users")

	u, err := s.User().Create(&model.User{
		Email: "testuser@example.com",
	})

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, tearDown := store.TestStore(t, databaseURL)
	defer tearDown()

	email := "testuser@example.com"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	s.User().Create(&model.User{
		Email: "testuser@example.com",
	})

	u, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
