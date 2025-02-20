package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Email:    "usertest@example.com",
		Password: "password",
	}
}
