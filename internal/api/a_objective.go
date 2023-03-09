package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/laidingqing/sokr/internal/ginx"
	"github.com/laidingqing/sokr/internal/schema"
	"github.com/laidingqing/sokr/internal/service"
)

type ObjectivesAPI struct {
	IObjectiveService service.IObjectiveService
}

func NewObjectivesAPI(objectiveService service.IObjectiveService) ObjectivesAPI {
	return ObjectivesAPI{IObjectiveService: objectiveService}
}

// Create 创建目标与关键结果
func (a *ObjectivesAPI) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var request schema.ReqestObjective
	if err := ginx.ParseJSON(c, &request); err != nil {
		ginx.ResError(c, err)
		return
	}

	var keyResults []*schema.KeyResult
	for _, v := range request.KeyResults {
		keyResults = append(keyResults, &schema.KeyResult{
			Name: v.Name,
		})
	}

	userID := ginx.GetUserID(c)
	objective := &schema.Objective{
		UserID:        userID,
		Name:          request.Name,
		ParentID:      request.ParentID,
		KeyResults:    keyResults,
		ObjectiveType: request.ObjectiveType,
		GroupID:       request.GroupID,
		CycleID:       request.CycleID,
	}

	result, err := a.IObjectiveService.Create(ctx, *objective)
	if err != nil {
		ginx.ResError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetAll 根据用户角色查询OKR
func (a *ObjectivesAPI) GetAll(c *gin.Context) {
	ctx := c.Request.Context()
	var request schema.ObjectiveQueryParam
	if err := ginx.ParseJSON(c, &request); err != nil {
		ginx.ResError(c, err)
		return
	}

	result, err := a.IObjectiveService.Query(ctx, request)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	// TODO ,wrap key_results of objective
	c.JSON(http.StatusOK, result)
}
