package service

import (
	"astroshot/gin-demo/pkg/service/dao"
	"astroshot/gin-demo/pkg/service/dao/model"
)

type UserService interface {
	Add(user *model.User) bool
	GetByID(id *int64) *model.User
}

type UserServiceImpl struct {
	UserDAO dao.UserDAO
}

func (service *UserServiceImpl) Add(user *model.User) bool {
	return service.UserDAO.Add(user)
}

func (service *UserServiceImpl) GetByID(id *int64) *model.User {
	return service.UserDAO.GetByID(id)
}
