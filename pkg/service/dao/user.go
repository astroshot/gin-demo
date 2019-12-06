package dao

import (
	"astroshot/gin-demo/pkg/service/bo"
	"astroshot/gin-demo/pkg/service/dao/model"
)

// UserDAO defines funcs to interfact with table `user`
type UserDAO interface {
	GetByID(id *int64) *model.User
	Add(user *model.User) bool
}

// UserDAOImpl implements interface UserDAO
type UserDAOImpl struct {
}

// GetByID returns User model by id
func (dao *UserDAOImpl) GetByID(id *int64) *model.User {
	user := model.User{}
	db.First(&user, *id)
	// or get a structure fulfilled with nil
	if user.ID == nil {
		return nil
	}

	return &user
}

// Add create User
func (dao *UserDAOImpl) Add(user *model.User) bool {
	if user == nil {
		return false
	}

	db.Create(&user)
	return true
}

// Update User
func (dao *UserDAOImpl) Update(user *model.User) bool {
	if user == nil {
		return false
	}

	db.Update(&user)
	return true
}

// GetByCondition returns Users
func (dao *UserDAOImpl) GetByCondition(condition *bo.UserQueryBO) {

}
