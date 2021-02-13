package service

import (
	"gin-demo/pkg/service/dao"
)

var (
	UserServiceInstance UserService
)

// InitService init package service
func InitService() {
	dao.InitDAO()
	UserServiceInstance = &UserServiceImpl{
		UserDAO: dao.UserDao,
	}
}
