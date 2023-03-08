package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/laidingqing/sokr/internal/ginx"
	"github.com/laidingqing/sokr/internal/schema"
	"github.com/laidingqing/sokr/internal/service"
)

type CycleAPI struct {
	ICycleService service.ICycleService
}

func NewCycleAPI(cycleService service.ICycleService) CycleAPI {
	return CycleAPI{ICycleService: cycleService}
}

func (a *CycleAPI) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.CreateCycleRequest
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}
	cycle := &schema.Cycle{
		Name:    item.Name,
		StartAt: item.StartAt,
		EndAt:   item.EndAt,
	}

	result, err := a.ICycleService.Create(ctx, *cycle)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	c.JSON(http.StatusOK, result)
}
