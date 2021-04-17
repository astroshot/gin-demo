package dao

import (
	// "fmt"
	"context"

	"gin-demo/pkg/config"
	"gin-demo/pkg/entity/bo"
	"gin-demo/pkg/entity/model"
)

// UserDAO defines funcs to interfact with table `user`
type UserDAO interface {
	GetByID(ctx context.Context, id *int64) *model.User
	Add(ctx context.Context, user *model.User) bool
	Update(ctx context.Context, user *model.User) bool
	GetByCondition(ctx context.Context, condition *bo.UserQueryBO) *bo.Pager
}

// UserDAOImpl implements interface UserDAO
type UserDAOImpl struct {
}

// GetByID returns User model by id
func (dao *UserDAOImpl) GetByID(ctx context.Context, id *int64) *model.User {
	user := model.User{}
	db.WithContext(ctx).First(&user, *id)
	logger := config.GetLoggerEntry(ctx)
	logger.Infof("Get User Result: %+v", user)
	// or get a structure fulfilled with nil
	if user.ID == nil {
		return nil
	}

	return &user
}

// Add create User
func (dao *UserDAOImpl) Add(ctx context.Context, user *model.User) bool {
	if user == nil {
		return false
	}

	if err := db.WithContext(ctx).Create(&user).Error; err != nil {
		panic(err)
	}
	return true
}

// Update User
func (dao *UserDAOImpl) Update(ctx context.Context, user *model.User) bool {
	if user == nil {
		return false
	}

	if err := db.WithContext(ctx).Model(&user).Updates(*user).Error; err != nil {
		panic(err)
	}
	return true
}

// GetByCondition returns Users
func (dao *UserDAOImpl) GetByCondition(ctx context.Context, condition *bo.UserQueryBO) *bo.Pager {
	var users []model.User
	var totalCount int64
	var totalCountInt int
	var pageCountInt int
	query := db.WithContext(ctx)

	if condition.Name != nil {
		query = query.Where("name LIKE ?", "%"+*condition.Name+"%")
	}

	if condition.PhoneNo != nil {
		query = query.Where("phone = ?", condition.PageNo)
	}

	offset := (*condition.PageNo - 1) * *condition.PageSize
	query.Find(&users).Count(&totalCount)
	query = query.Limit(*condition.PageSize)
	query = query.Offset(offset)
	query.Find(&users)

	totalCountInt = int(totalCount)
	pageCountInt = (totalCountInt + *condition.PageSize - 1) / *condition.PageSize

	pager := &bo.Pager{
		PageNo:     condition.PageNo,
		PageSize:   condition.PageSize,
		PageCount:  &pageCountInt,
		TotalCount: &totalCountInt,
		Data:       users,
	}

	return pager
}
