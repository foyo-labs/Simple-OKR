package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/laidingqing/sokr/internal/ginx"
	"github.com/laidingqing/sokr/internal/schema"
	"github.com/laidingqing/sokr/internal/service"
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

	a.IUserService.Verify(ctx, item.Email, item.Password)

	c.JSON(http.StatusOK, "Ok")
}
