package dao

import (
	"astroshot/gin-demo/pkg/model"
)

// UserDAO defines funcs to interfact with table `user`
type UserDAO interface {
	GetById(id *int64) *model.User
	Add(user *model.User) bool
}
