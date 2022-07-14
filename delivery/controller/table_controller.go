package controller

import (
	"go-api-with-gin2/delivery/api"
	"go-api-with-gin2/model"
	"go-api-with-gin2/usecase"
	"go-api-with-gin2/utils"

	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TableController struct {
	router        *gin.RouterGroup
	ucTable       usecase.CreateTableUseCase
	ucTableList   usecase.ListTableUseCase
	ucTableDelete usecase.DeleteTableUseCase
	ucTableUpdate usecase.UpdateTableUseCase
	api.BaseApi
}

func (b *TableController) createTable(c *gin.Context) {
	var newTable model.Table
	err := b.ParseRequestBody(c, &newTable)
	if err != nil {
		b.Failed(c, utils.RequiredError())
		return
	}
	err = b.ucTable.CreateTable(&newTable)
	if err != nil {
		b.Failed(c, err)
		return
	}
	b.Success(c, newTable)
}

func (b *TableController) findAllTable(c *gin.Context) {
	Tables, err := b.ucTableList.Retrive()
	if err != nil {
		b.Failed(c, err)
		return
	}
	b.Success(c, Tables)
}

func (m *TableController) deleteTable(c *gin.Context) {
	tableId, _ := strconv.Atoi(c.Query("id"))
	delMenu := model.Table{Model: gorm.Model{ID: uint(tableId)}}
	err := m.ucTableDelete.DeleteTable(&delMenu)
	if err != nil {
		m.Failed(c, err)
		return
	}
	m.Success(c, tableId)
}

func (m *TableController) updateTable(c *gin.Context) {
	var newTable model.Table
	err := m.ParseRequestBody(c, &newTable)
	by := map[string]interface{}{
		"TableDescription": newTable.TableDescription,
		"IsAvailable":      newTable.IsAvailable,
	}

	err = m.ucTableUpdate.UpdateTable(&newTable, by)
	if err != nil {
		m.Failed(c, err)
		return
	}
	m.Success(c, newTable)
}

func NewTableController(
	router *gin.RouterGroup,
	ucTable usecase.CreateTableUseCase,
	ucTableList usecase.ListTableUseCase,
	ucTableDelete usecase.DeleteTableUseCase,
	ucTableUpdate usecase.UpdateTableUseCase) *TableController {
	var controller TableController = TableController{
		router:        router,
		ucTable:       ucTable,
		ucTableList:   ucTableList,
		ucTableDelete: ucTableDelete,
		ucTableUpdate: ucTableUpdate,
	}

	router.POST("/table", controller.createTable)
	router.GET("/table", controller.findAllTable)
	router.DELETE("/table/delete", controller.deleteTable)
	router.PATCH("/table/update", controller.updateTable)
	return &controller
}
