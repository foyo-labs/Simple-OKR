package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/laidingqing/sokr/internal/ginx"
	"github.com/laidingqing/sokr/internal/schema"
	"github.com/laidingqing/sokr/internal/service"
	"github.com/laidingqing/sokr/pkg/errors"
	"github.com/laidingqing/sokr/pkg/logger"
)

type UserAPI struct {
	IUserService service.IUserService
}

func NewUserAPI(userService service.IUserService) UserAPI {
	return UserAPI{IUserService: userService}
}

func (a *UserAPI) Login(c *gin.Context) {

	ctx := c.Request.Context()
	var item schema.LoginParam
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	result, err := a.IUserService.Verify(ctx, item.Email, item.Password)
	if err != nil {
		ginx.ResError(c, err)
		return
	}

	ginx.SetUserID(c, result.ID)
	tokenInfo, err := a.IUserService.GenerateToken(ctx, result.ID)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	c.JSON(http.StatusOK, tokenInfo)
}

// Create 创建用户，
func (a *UserAPI) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.ReqestRegistion
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	} else if item.Password == "" {
		ginx.ResError(c, errors.New400Response("密码不能为空"))
		return
	}

	var user schema.User
	user.Email = item.Email
	user.Password = item.Password

	result, err := a.IUserService.Create(ctx, user)
	if err != nil {
		logger.Errorf("%v", err)
		ginx.ResError(c, err)
		return
	}

	ginx.ResSuccess(c, result)
}