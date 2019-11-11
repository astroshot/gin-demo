package impl

import (
	"astroshot/gin-demo/pkg/model"
)

// UserDAOImpl implements interface UserDAO
type UserDAOImpl struct {
	BaseDAOImpl
}

// GetByID returns User model by id
func (dao UserDAOImpl) GetByID(id *int64) *model.User {
	user := model.User{}
	dao.db.First(&user, *id)
	return &user
}

func (dao UserDAOImpl) Add(user *model.User) bool {
	if user == nil {
		return false
	}

	dao.db.Create(&user)
	return true
}
