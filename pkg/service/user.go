package service

import (
	"context"

	"gin-demo/pkg/service/bo"
	"gin-demo/pkg/service/dao"
	"gin-demo/pkg/service/dao/model"
)

type UserService interface {
	Add(ctx context.Context, user *model.User) bool
	Update(ctx context.Context, user *model.User) bool
	GetByID(ctx context.Context, id *int64) *model.User
	GetByCondition(ctx context.Context, condition *bo.UserQueryBO) *bo.Pager
}

type UserServiceImpl struct {
	UserDAO dao.UserDAO
}

func (service *UserServiceImpl) Add(ctx context.Context, user *model.User) bool {
	return service.UserDAO.Add(ctx, user)
}

func (service *UserServiceImpl) GetByID(ctx context.Context, id *int64) *model.User {
	return service.UserDAO.GetByID(ctx, id)
}

func (service *UserServiceImpl) GetByCondition(ctx context.Context, condition *bo.UserQueryBO) *bo.Pager {
	return service.UserDAO.GetByCondition(ctx, condition)
}

func (service *UserServiceImpl) Update(ctx context.Context, user *model.User) bool {
	return service.UserDAO.Update(ctx, user)
}
