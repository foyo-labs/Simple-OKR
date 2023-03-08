package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/laidingqing/sokr/internal/ginx"
	"github.com/laidingqing/sokr/internal/schema"
	"github.com/laidingqing/sokr/internal/service"
)

type GroupAPI struct {
	IGroupService service.IGroupService
}

func NewGroupAPI(groupService service.IGroupService) GroupAPI {
	return GroupAPI{IGroupService: groupService}
}

func (a *GroupAPI) GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, "Ok")
}

// Create 创建组数据
func (a *GroupAPI) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.CreateGroupRequest
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}
	userID := ginx.GetUserID(c)
	group := &schema.Group{
		UserID:   userID,
		Name:     item.Name,
		ParentID: item.ParentID,
	}

	result, err := a.IGroupService.Create(ctx, *group)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	c.JSON(http.StatusOK, result)
}
