package controller

import (
	"go-api-with-gin2/delivery/api"
	"go-api-with-gin2/model"
	"go-api-with-gin2/usecase"
	"go-api-with-gin2/utils"

	"github.com/gin-gonic/gin"
)

type BillController struct {
	router     *gin.RouterGroup
	ucBill     usecase.CreateBillUseCase
	ucBillList usecase.ListBillUseCase
	api.BaseApi
}

func (b *BillController) createBill(c *gin.Context) {
	var newBill model.Bill
	err := b.ParseRequestBody(c, &newBill)
	if err != nil {
		b.Failed(c, utils.RequiredError())
		return
	}
	err = b.ucBill.CreateBill(&newBill)
	if err != nil {
		b.Failed(c, err)
		return
	}
	b.Success(c, newBill)
}

func (b *BillController) findAllBill(c *gin.Context) {
	bills, err := b.ucBillList.Retrive()
	if err != nil {
		b.Failed(c, err)
		return
	}
	b.Success(c, bills)
}

func NewBillController(
	router *gin.RouterGroup,
	ucBill usecase.CreateBillUseCase,
	ucBillList usecase.ListBillUseCase) *BillController {
	var controller BillController = BillController{
		router:     router,
		ucBill:     ucBill,
		ucBillList: ucBillList,
	}

	router.POST("/bill", controller.createBill)
	router.GET("/bill", controller.findAllBill)
	return &controller
}
