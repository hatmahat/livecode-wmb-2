package controller

import (
	"go-api-with-gin2/delivery/api"
	"go-api-with-gin2/model"
	"go-api-with-gin2/usecase"
	"go-api-with-gin2/utils"

	"github.com/gin-gonic/gin"
)

type BillDetailController struct {
	router           *gin.RouterGroup
	ucBillDetail     usecase.CreateBillDetailUseCase
	ucBillDetailList usecase.ListBillDetailUseCase
	api.BaseApi
}

func (b *BillDetailController) createBillDetail(c *gin.Context) {
	var newBill model.BillDetail
	err := b.ParseRequestBody(c, &newBill)
	if err != nil {
		b.Failed(c, utils.RequiredError())
		return
	}
	err = b.ucBillDetail.CreateBillDetail(&newBill)
	if err != nil {
		b.Failed(c, err)
		return
	}
	b.Success(c, newBill)
}

func (b *BillDetailController) findAllBillDetail(c *gin.Context) {
	bills, err := b.ucBillDetailList.Retrive()
	if err != nil {
		b.Failed(c, err)
		return
	}
	b.Success(c, bills)
}

func NewBillDetailController(
	router *gin.RouterGroup,
	ucBillDetail usecase.CreateBillDetailUseCase,
	ucBillDetailList usecase.ListBillDetailUseCase) *BillDetailController {
	var controller BillDetailController = BillDetailController{
		router:           router,
		ucBillDetail:     ucBillDetail,
		ucBillDetailList: ucBillDetailList,
	}

	router.POST("/bill_detail", controller.createBillDetail)
	router.GET("/bill_detail", controller.findAllBillDetail)
	return &controller
}
