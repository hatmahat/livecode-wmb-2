package controller

import (
	"go-api-with-gin2/delivery/api"
	"go-api-with-gin2/model"
	"go-api-with-gin2/usecase"
	"go-api-with-gin2/utils"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	router           *gin.RouterGroup
	ucCustomer       usecase.CreateCustomerUseCase
	ucCustomerList   usecase.ListCustomerUseCase
	ucCustomerUpdate usecase.UpdateCustomerUseCase
	api.BaseApi
}

func (b *CustomerController) createCustomer(c *gin.Context) {
	var newCustomer model.Customer
	err := b.ParseRequestBody(c, &newCustomer)
	if err != nil {
		b.Failed(c, utils.RequiredError())
		return
	}
	err = b.ucCustomer.CreateCustomer(&newCustomer)
	if err != nil {
		b.Failed(c, err)
		return
	}
	b.Success(c, newCustomer)
}

func (b *CustomerController) findAllCustomer(c *gin.Context) {
	Customers, err := b.ucCustomerList.Retrive()
	if err != nil {
		b.Failed(c, err)
		return
	}
	b.Success(c, Customers)
}

func (m *CustomerController) updateCustomer(c *gin.Context) {
	var newCustomer model.Customer
	err := m.ParseRequestBody(c, &newCustomer)
	by := map[string]interface{}{
		"IsMember":  newCustomer.IsMember,
		"Discounts": newCustomer.Discounts,
	}
	err = m.ucCustomerUpdate.UpdateCustomer(&newCustomer, by)
	if err != nil {
		m.Failed(c, err)
		return
	}
	m.Success(c, newCustomer)
}

func NewCustomerController(
	router *gin.RouterGroup,
	ucCustomer usecase.CreateCustomerUseCase,
	ucCustomerList usecase.ListCustomerUseCase,
	ucCustomerUpdate usecase.UpdateCustomerUseCase) *CustomerController {
	var controller CustomerController = CustomerController{
		router:           router,
		ucCustomer:       ucCustomer,
		ucCustomerList:   ucCustomerList,
		ucCustomerUpdate: ucCustomerUpdate,
	}

	router.POST("/customer", controller.createCustomer)
	router.GET("/customer", controller.findAllCustomer)
	router.PATCH("/customer/register", controller.updateCustomer)
	return &controller
}
