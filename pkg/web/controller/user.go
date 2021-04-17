package controller

import (
	view "gin-demo/pkg/common/model"
	"gin-demo/pkg/config"
	"gin-demo/pkg/entity/bo"
	"gin-demo/pkg/entity/model"
	"gin-demo/pkg/entity/vo"
	"gin-demo/pkg/service"
	"gin-demo/pkg/util"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ListUsers returns list of users
func ListUsers(c *gin.Context) {
	name := util.GetQueryStr(c, "name")
	phoneNo := util.GetQueryStr(c, "phoneNo")
	pageNo := util.GetQueryInt(c, "pageNo")
	pageSize := util.GetQueryInt(c, "pageSize")

	query := &bo.UserQueryBO{
		Name:     name,
		PhoneNo:  phoneNo,
		PageNo:   pageNo,
		PageSize: pageSize,
	}

	pager := service.UserServiceInstance.GetByCondition(c, query)
	res := view.Success(0, util.SuccessInfo, pager)
	c.JSON(http.StatusOK, res)
}

// AddUser creates model User in db
func AddUser(c *gin.Context) {
	var userVO vo.UserVO
	var res *view.JSONResponse
	if err := c.ShouldBindJSON(&userVO); err != nil {
		res = view.Fail(-1, util.FailInfo, err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}

	status := util.StatusNormal
	user := &model.User{
		Name:        userVO.Name,
		Email:       userVO.Email,
		Phone:       userVO.Phone,
		Description: userVO.Description,
		Status:      &status,
	}
	service.UserServiceInstance.Add(c, user)
	// Logger.Infof("Log: %s", user.Name)
	res = view.Success(0, util.SuccessInfo, true)
	c.JSON(http.StatusOK, res)
}

func UpdateUser(c *gin.Context) {
	var userVO vo.UserVO
	var res *view.JSONResponse

	userID := c.Param("token")
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		res = view.Fail(-1, util.FailInfo, nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := c.ShouldBindJSON(&userVO); err != nil {
		res = view.Fail(-1, util.FailInfo, err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}

	user := service.UserServiceInstance.GetByID(c, &id)
	if user == nil {
		res = view.Fail(-1, "user not found", nil)
		c.JSON(http.StatusNotFound, res)
		return
	}
	user.Name = userVO.Name
	user.Email = userVO.Email
	user.Description = userVO.Description
	user.Phone = userVO.Phone

	service.UserServiceInstance.Update(c, user)
	res = view.Success(0, util.SuccessInfo, true)
	c.JSON(http.StatusOK, res)
}

// GetUserByID returns User by id
func GetUserByID(c *gin.Context) {
	var res *view.JSONResponse
	logger := config.GetLoggerEntry(c)
	userID := c.Param("token")
	logger.WithContext(c).Infof("Request By id: %s", userID)
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		res = view.Fail(-1, util.FailInfo, nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	user := service.UserServiceInstance.GetByID(c, &id)
	if user == nil {
		res = view.Fail(-1, "user not found", nil)
		c.JSON(http.StatusNotFound, res)
		return
	}

	res = view.Success(0, util.SuccessInfo, user)
	c.JSON(http.StatusOK, res)
}
