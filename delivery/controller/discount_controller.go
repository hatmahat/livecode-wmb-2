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

type DiscountController struct {
	router           *gin.RouterGroup
	ucDiscount       usecase.CreateDiscountUseCase
	ucDiscountList   usecase.ListDiscountUseCase
	ucDiscountDelete usecase.DeleteDiscountUseCase
	ucDiscountUpdate usecase.UpdateDiscountUseCase
	api.BaseApi
}

func (b *DiscountController) createDiscount(c *gin.Context) {
	var newDiscount model.Discount
	err := b.ParseRequestBody(c, &newDiscount)
	if err != nil {
		b.Failed(c, utils.RequiredError())
		return
	}
	err = b.ucDiscount.CreateDiscount(&newDiscount)
	if err != nil {
		b.Failed(c, err)
		return
	}
	b.Success(c, newDiscount)
}

func (b *DiscountController) findAllDiscount(c *gin.Context) {
	Discounts, err := b.ucDiscountList.Retrive()
	if err != nil {
		b.Failed(c, err)
		return
	}
	b.Success(c, Discounts)
}

func (m *DiscountController) deleteDiscount(c *gin.Context) {
	discountId, _ := strconv.Atoi(c.Query("id"))
	delDiscount := model.Discount{Model: gorm.Model{ID: uint(discountId)}}
	err := m.ucDiscountDelete.DeleteDiscount(&delDiscount)
	if err != nil {
		m.Failed(c, err)
		return
	}
	m.Success(c, discountId)
}

func (m *DiscountController) updateDiscount(c *gin.Context) {
	var newDiscount model.Discount
	err := m.ParseRequestBody(c, &newDiscount)
	by := map[string]interface{}{
		"Description": newDiscount.Description,
		"Pct":         newDiscount.Pct,
		//"Customer":    newDiscount.Customers,
	}
	err = m.ucDiscountUpdate.UpdateDiscount(&newDiscount, by)
	if err != nil {
		m.Failed(c, err)
		return
	}
	m.Success(c, newDiscount)
}

func NewDiscountController(
	router *gin.RouterGroup,
	ucDiscount usecase.CreateDiscountUseCase,
	ucDiscountList usecase.ListDiscountUseCase,
	ucDiscountDelete usecase.DeleteDiscountUseCase,
	ucDiscountUpdate usecase.UpdateDiscountUseCase) *DiscountController {
	var controller DiscountController = DiscountController{
		router:           router,
		ucDiscount:       ucDiscount,
		ucDiscountList:   ucDiscountList,
		ucDiscountDelete: ucDiscountDelete,
		ucDiscountUpdate: ucDiscountUpdate,
	}

	router.POST("/discount", controller.createDiscount)
	router.GET("/discount", controller.findAllDiscount)
	router.DELETE("/discount/delete", controller.deleteDiscount)
	router.PATCH("/discount/update", controller.updateDiscount)
	return &controller
}
