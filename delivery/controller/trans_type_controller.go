package controller

import (
	"go-api-with-gin2/delivery/api"
	"go-api-with-gin2/model"
	"go-api-with-gin2/usecase"
	"go-api-with-gin2/utils"

	"github.com/gin-gonic/gin"
)

type TransTypeController struct {
	router            *gin.RouterGroup
	ucTransType       usecase.CreateTransTypeUseCase
	ucTransTypeList   usecase.ListTransTypeUseCase
	ucTransTypeDelete usecase.DeleteTransTypeUseCase
	ucTransTypeUpdate usecase.UpdateTransTypeUseCase
	api.BaseApi
}

func (b *TransTypeController) createTransType(c *gin.Context) {
	var newTransType model.TransType
	err := b.ParseRequestBody(c, &newTransType)
	if err != nil {
		b.Failed(c, utils.RequiredError())
		return
	}
	err = b.ucTransType.CreateTransType(&newTransType)
	if err != nil {
		b.Failed(c, err)
		return
	}
	b.Success(c, newTransType)
}

func (b *TransTypeController) findAllTransType(c *gin.Context) {
	TransTypes, err := b.ucTransTypeList.Retrive()
	if err != nil {
		b.Failed(c, err)
		return
	}
	b.Success(c, TransTypes)
}

func (m *TransTypeController) deleteTransType(c *gin.Context) {
	transTypeID := c.Query("trans_type_id")
	delTrans := model.TransType{Id: transTypeID}
	err := m.ucTransTypeDelete.DeleteTransType(&delTrans)
	if err != nil {
		m.Failed(c, err)
		return
	}
	m.Success(c, transTypeID)
}

func (m *TransTypeController) updateTransType(c *gin.Context) {
	var newTransType model.TransType
	err := m.ParseRequestBody(c, &newTransType)
	by := map[string]interface{}{
		"Description": newTransType.Description,
	}
	err = m.ucTransTypeUpdate.UpdateTransType(&newTransType, by)
	if err != nil {
		m.Failed(c, err)
		return
	}
	m.Success(c, newTransType)
}

func NewTransTypeController(
	router *gin.RouterGroup,
	ucTransType usecase.CreateTransTypeUseCase,
	ucTransTypeList usecase.ListTransTypeUseCase,
	ucTransTypeDelete usecase.DeleteTransTypeUseCase,
	ucTransTypeUpdate usecase.UpdateTransTypeUseCase) *TransTypeController {
	var controller TransTypeController = TransTypeController{
		router:            router,
		ucTransType:       ucTransType,
		ucTransTypeList:   ucTransTypeList,
		ucTransTypeDelete: ucTransTypeDelete,
		ucTransTypeUpdate: ucTransTypeUpdate,
	}

	router.POST("/trans_type", controller.createTransType)
	router.GET("/trans_type", controller.findAllTransType)
	router.DELETE("/trans_type/delete", controller.deleteTransType)
	router.PATCH("/trans_type/update", controller.updateTransType)
	return &controller
}
