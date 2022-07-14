package controller

import (
	"go-api-with-gin2/delivery/api"
	"go-api-with-gin2/model"
	"go-api-with-gin2/usecase"
	"go-api-with-gin2/utils"

	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
	// "strconv"
)

type MenuController struct {
	router       *gin.RouterGroup
	ucMenu       usecase.CreateMenuUseCase
	UcMenuList   usecase.ListMenuUseCase
	ucMenuDelete usecase.DeleteMenuUseCase
	ucMenuUpdate usecase.UpdateMenuUseCase
	api.BaseApi
}

func (m *MenuController) createMenu(c *gin.Context) {
	var newMenu model.Menu
	err := m.ParseRequestBody(c, &newMenu)
	if err != nil {
		m.Failed(c, utils.RequiredError())
		return
	}
	err = m.ucMenu.CreateMenu(&newMenu)
	if err != nil {
		m.Failed(c, err)
		return
	}
	m.Success(c, newMenu)
}

func (m *MenuController) listMenu(c *gin.Context) {
	menus, err := m.UcMenuList.ListMenu()
	if err != nil {
		m.Failed(c, err)
		return
	}
	m.Success(c, menus)
}

func (m *MenuController) deleteMenu(c *gin.Context) {
	menuName := c.Query("menu_name")
	delMenu := model.Menu{MenuName: menuName}
	err := m.ucMenuDelete.DeleteMenu(&delMenu)
	if err != nil {
		m.Failed(c, err)
		return
	}
	m.Success(c, menuName)
}

func (m *MenuController) updateMenu(c *gin.Context) {
	oldMenuName := c.Query("old_menu_name")
	newMenuName := c.Query("new_menu_name")
	var newMenu model.Menu = model.Menu{
		MenuName: oldMenuName,
	}
	by := map[string]interface{}{
		"MenuName": newMenuName,
	}
	err := m.ucMenuUpdate.UpdateMenu(&newMenu, by)
	if err != nil {
		m.Failed(c, err)
		return
	}
	m.Success(c, newMenu)
}

func NewMenuController(
	router *gin.RouterGroup,
	ucMenu usecase.CreateMenuUseCase,
	UcMenuList usecase.ListMenuUseCase,
	ucMenuDelete usecase.DeleteMenuUseCase,
	ucMenuUpdate usecase.UpdateMenuUseCase) *MenuController {
	var controller MenuController = MenuController{
		router:       router,
		ucMenu:       ucMenu,
		UcMenuList:   UcMenuList,
		ucMenuDelete: ucMenuDelete,
		ucMenuUpdate: ucMenuUpdate,
	}

	router.POST("/menu", controller.createMenu)
	router.GET("/menu", controller.listMenu)
	router.DELETE("/menu/delete", controller.deleteMenu)
	router.PATCH("/menu/update", controller.updateMenu)
	return &controller
}
