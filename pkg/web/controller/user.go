package controller

import (
	view "astroshot/gin-demo/pkg/common/model"
	"astroshot/gin-demo/pkg/service"
	dao_model "astroshot/gin-demo/pkg/service/dao/model"
	"astroshot/gin-demo/pkg/util"
	"astroshot/gin-demo/pkg/web/model"

	"net/http"

	"github.com/gin-gonic/gin"
)

// ListUsers returns list of users
func ListUsers(c *gin.Context) {

}

// AddUser creates model User in db
func AddUser(c *gin.Context) {
	var userVO model.UserVO
	var res *view.JSONResponse
	if err := c.ShouldBindJSON(&userVO); err != nil {
		res = view.Fail(-1, util.FailInfo, err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}

	status := util.StatusNormal
	user := &dao_model.User{
		Name:        userVO.Name,
		Email:       userVO.Email,
		Phone:       userVO.Phone,
		Description: userVO.Description,
		Status:      &status,
	}
	service.UserServiceInstance.Add(user)
	// Logger.Infof("Log: %s", user.Name)
	res = view.Success(0, util.SuccessInfo, true)
	c.JSON(http.StatusOK, res)
}
