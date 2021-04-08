package service

import (
	"context"

	"gin-demo/pkg/service/bo"
	"gin-demo/pkg/service/dao"
	"gin-demo/pkg/service/dao/model"
)

type UserService interface {
	Add(ctx context.Context, user *model.User) bool
	Update(user *model.User) bool
	GetByID(id *int64) *model.User
	GetByCondition(condition *bo.UserQueryBO) *bo.Pager
}

type UserServiceImpl struct {
	UserDAO dao.UserDAO
}

func (service *UserServiceImpl) Add(ctx context.Context, user *model.User) bool {
	return service.UserDAO.Add(ctx, user)
}

func (service *UserServiceImpl) GetByID(id *int64) *model.User {
	return service.UserDAO.GetByID(id)
}

func (service *UserServiceImpl) GetByCondition(condition *bo.UserQueryBO) *bo.Pager {
	return service.UserDAO.GetByCondition(condition)
}

func (service *UserServiceImpl) Update(user *model.User) bool {
	return service.UserDAO.Update(user)
}
