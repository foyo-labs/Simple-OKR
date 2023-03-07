package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/laidingqing/sokr/internal/service"
)

type UnitAPI struct {
	IUnitService service.IUnitService
}

func NewUnitAPI(unitService service.IUnitService) UnitAPI {
	return UnitAPI{IUnitService: unitService}
}

func (a *UnitAPI) GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, "Ok")
}

func (a *UnitAPI) CreateCompany(c *gin.Context) {
	c.JSON(http.StatusOK, "Ok")
}
