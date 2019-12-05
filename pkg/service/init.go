package service

import (
	"astroshot/gin-demo/pkg/service/dao"
)

var (
	UserServiceInstance UserService
)

func InitService() {
	dao.InitDAO()
	UserServiceInstance = &UserServiceImpl{
		UserDAO: dao.UserDao,
	}
}
