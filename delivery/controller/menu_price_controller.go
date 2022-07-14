package controller

import (
	"go-api-with-gin2/delivery/api"
	"go-api-with-gin2/model"
	"go-api-with-gin2/usecase"
	"go-api-with-gin2/utils"

	"strconv"

	"github.com/gin-gonic/gin"
)

type MenuPriceController struct {
	router            *gin.RouterGroup
	ucMenuPrice       usecase.CreateMenuPriceUseCase
	ucMenuPriceList   usecase.ListMenuPriceUseCase
	ucMenuPriceDelete usecase.DeleteMenuPriceUseCase
	ucMenuPriceUpdate usecase.UpdateMenuPriceUseCase
	api.BaseApi
}

func (m *MenuPriceController) createNewMenuPrice(c *gin.Context) {
	var newMenu model.MenuPrice
	err := m.ParseRequestBody(c, &newMenu)
	if err != nil {
		m.Failed(c, utils.RequiredError())
		return
	}
	err = m.ucMenuPrice.CreateMenuPrice(&newMenu)
	if err != nil {
		m.Failed(c, err)
		return
	}
	m.Success(c, newMenu)
}

func (m *MenuPriceController) findAllMenuPrice(c *gin.Context) {
	menus, err := m.ucMenuPriceList.Retrive()
	if err != nil {
		m.Failed(c, err)
		return
	}
	m.Success(c, menus)
}

func (m *MenuPriceController) deleteMenuPrice(c *gin.Context) {
	menuID, _ := strconv.Atoi(c.Query("menu_id"))
	delMenu := model.MenuPrice{MenuID: uint(menuID)}
	err := m.ucMenuPriceDelete.DeleteMenuPrice(&delMenu)
	if err != nil {
		m.Failed(c, err)
		return
	}
	m.Success(c, menuID)
}

func (m *MenuPriceController) updateMenuPrice(c *gin.Context) {
	//menuID, _ := strconv.Atoi(c.Query("menu_id"))
	var newMenu model.MenuPrice
	err := m.ParseRequestBody(c, &newMenu)
	by := map[string]interface{}{
		//"Menu":  newMenu.Menu,
		"Price": newMenu.Price,
	}
	//menuId := model.MenuPrice{MenuID: newMenu.MenuID}
	err = m.ucMenuPriceUpdate.UpdateMenuPrice(&newMenu, by)
	if err != nil {
		m.Failed(c, err)
		return
	}
	m.Success(c, newMenu)
}

func NewMenuPriceController(
	router *gin.RouterGroup,
	ucMenuPrice usecase.CreateMenuPriceUseCase,
	ucMenuPriceList usecase.ListMenuPriceUseCase,
	ucMenuPriceDelete usecase.DeleteMenuPriceUseCase,
	ucMenuPriceUpdate usecase.UpdateMenuPriceUseCase) *MenuPriceController {
	var controller MenuPriceController = MenuPriceController{
		router:            router,
		ucMenuPrice:       ucMenuPrice,
		ucMenuPriceList:   ucMenuPriceList,
		ucMenuPriceDelete: ucMenuPriceDelete,
		ucMenuPriceUpdate: ucMenuPriceUpdate,
	}

	router.POST("/menu_price", controller.createNewMenuPrice)
	router.GET("/menu_price", controller.findAllMenuPrice)
	router.DELETE("/menu_price/delete", controller.deleteMenuPrice)
	router.PATCH("/menu_price/update", controller.updateMenuPrice)
	return &controller
}
