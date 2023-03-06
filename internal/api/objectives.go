package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/laidingqing/sokr/internal/service"
)

type ObjectivesAPI struct {
	IObjectiveService service.IObjectiveService
}

func NewObjectivesAPI(objectiveService service.IObjectiveService) ObjectivesAPI {
	return ObjectivesAPI{IObjectiveService: objectiveService}
}

func (a *ObjectivesAPI) GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, "Ok")
}
