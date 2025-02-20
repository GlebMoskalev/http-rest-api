package store

import "github.com/GlebMoskalev/http-rest-api/internal/app/model"

type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(string) (*model.User, error)
}
