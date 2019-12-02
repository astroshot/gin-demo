package service

import (
	"astroshot/gin-demo/pkg/service/dao"
	"astroshot/gin-demo/pkg/service/dao/model"
)

type UserService interface {
	Add(user *model.User) bool
}

type UserServiceImpl struct {
	UserDAO *dao.UserDAOImpl
}

func (service *UserServiceImpl) Add(user *model.User) bool {
	return service.UserDAO.Add(user)
}

func (service *UserServiceImpl) GetById(id *int) *model.User {
	return service.GetById(id)
}
