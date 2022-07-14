package controller

import (
	"go-api-with-gin2/delivery/api"
	"go-api-with-gin2/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
	// "strconv"
)

type LopeiController struct {
	router       *gin.RouterGroup
	lopeiUsecase usecase.CheckBalanceUseCaseLopei
	api.BaseApi
}

func (m *LopeiController) checkBalance(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	// var newMenu model.Menu = model.Menu{
	// 	MenuName: oldMenuName,
	// }
	// by := map[string]interface{}{
	// 	"MenuName": newMenuName,
	// }
	balance, err := m.lopeiUsecase.GetBalance(int32(id))
	if err != nil {
		m.Failed(c, err)
		return
	}
	m.Success(c, balance)
}

func NewLopeiController(
	router *gin.RouterGroup,
	lopeiUsecase usecase.CheckBalanceUseCaseLopei) *LopeiController {
	var controller LopeiController = LopeiController{
		router:       router,
		lopeiUsecase: lopeiUsecase,
	}
	router.GET("/lopei", controller.checkBalance)
	return &controller
}
