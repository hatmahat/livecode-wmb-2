package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go-api-with-gin2/delivery/api/dto"
	"go-api-with-gin2/model"
	"go-api-with-gin2/usecase"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CreateMenuUseCaseMock struct{ mock.Mock }

type ListMenuUseCaseMock struct{ mock.Mock }

type DeleteMenuUseCaseMock struct{ mock.Mock }

type UpdateMenuUseCaseMock struct{ mock.Mock }

type MenuControllerTestSuite struct {
	suite.Suite
	routerMock            *gin.Engine
	routerGroupMock       *gin.RouterGroup
	createMenuUseCaseMock *CreateMenuUseCaseMock
	listMenuUseCaseMock   *ListMenuUseCaseMock
	// deleteMenuUseCaseMock *DeleteMenuUseCaseMock
	// updateMenuUseCaseMock *UpdateMenuUseCaseMock
	deleteMenuUseCaseMock usecase.DeleteMenuUseCase
	updateMenuUseCaseMock usecase.UpdateMenuUseCase
}

var dummyMenus = []model.Menu{
	{
		MenuName: "Ayam Greprek",
	},
}

func (suite *MenuControllerTestSuite) SetupTest() {
	suite.routerMock = gin.Default()
	suite.routerGroupMock = suite.routerMock.Group("/api/master")
	suite.createMenuUseCaseMock = new(CreateMenuUseCaseMock)
	suite.listMenuUseCaseMock = new(ListMenuUseCaseMock)
	// suite.deleteMenuUseCaseMock = new(DeleteMenuUseCaseMock)
	// suite.updateMenuUseCaseMock = new(UpdateMenuUseCaseMock)
}

func (m *CreateMenuUseCaseMock) CreateMenu(menu *model.Menu) error {
	args := m.Called(*menu)
	if args.Get(0) != nil {
		return args.Get(0).(error)
	}
	return nil
}

func (suite *MenuControllerTestSuite) TestCreateMenuApi_Success() {
	dummyMenu := dummyMenus[0]
	suite.createMenuUseCaseMock.On("CreateMenu", dummyMenu).Return(nil)
	NewMenuController(
		suite.routerGroupMock,
		suite.createMenuUseCaseMock,
		suite.listMenuUseCaseMock,
		suite.deleteMenuUseCaseMock,
		suite.updateMenuUseCaseMock,
	)
	r := httptest.NewRecorder()
	reqBody, _ := json.Marshal(dummyMenu)
	request, _ := http.NewRequest(http.MethodPost, "/api/master/menu", bytes.NewBuffer(reqBody))
	suite.routerMock.ServeHTTP(r, request)
	response := r.Body.String()
	var actualResponse dto.MenuResponse
	json.Unmarshal([]byte(response), &actualResponse)
	fmt.Println("RESPONSE", actualResponse)
	actualMenu := actualResponse.Data
	assert.Equal(suite.T(), http.StatusOK, r.Code)
	assert.Equal(suite.T(), dummyMenu.MenuName, actualMenu.MenuName)
}

func (suite *MenuControllerTestSuite) TestCreateMenuApi_Failed() {
	dummyMenu := dummyMenus[0]
	suite.createMenuUseCaseMock.On("CreateMenu", dummyMenu).Return(errors.New("failed"))
	NewMenuController(
		suite.routerGroupMock,
		suite.createMenuUseCaseMock,
		suite.listMenuUseCaseMock,
		suite.deleteMenuUseCaseMock,
		suite.updateMenuUseCaseMock,
	)
	r := httptest.NewRecorder()
	reqBody, _ := json.Marshal(dummyMenu)
	request, _ := http.NewRequest(http.MethodPost, "/api/master/menu", bytes.NewBuffer(reqBody))
	suite.routerMock.ServeHTTP(r, request)
	assert.Equal(suite.T(), http.StatusInternalServerError, r.Code)
	response := r.Body.String()
	fmt.Println("ERROR RESPONSE", response)
	var errorResponse dto.ErrorCode
	json.Unmarshal([]byte(response), &errorResponse)
	fmt.Println("ERROR RESPONSE", errorResponse.Code)

	assert.Equal(suite.T(), "XX", errorResponse.Code)
}

func (c *ListMenuUseCaseMock) ListMenu() ([]model.Menu, error) {
	args := c.Called()
	if args.Get(1) != nil {
		return nil, args.Get(1).(error)
	}
	return args.Get(0).([]model.Menu), nil
}

func (suite *MenuControllerTestSuite) TestListMenuApi_Success() {
	dummyMenu := dummyMenus
	suite.listMenuUseCaseMock.On("ListMenu").Return(dummyMenu, nil)
	NewMenuController(
		suite.routerGroupMock,
		suite.createMenuUseCaseMock,
		suite.listMenuUseCaseMock,
		suite.deleteMenuUseCaseMock,
		suite.updateMenuUseCaseMock,
	)
	r := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/api/master/menu", nil)
	suite.routerMock.ServeHTTP(r, request)
	var actualResponse dto.ListMenuResponse
	response := r.Body.String()
	json.Unmarshal([]byte(response), &actualResponse)
	fmt.Println("ACTUAL RESPONSE", actualResponse)
	actualMenu := actualResponse.Data
	fmt.Println("ACTUAL MENU", actualMenu)
	assert.Equal(suite.T(), http.StatusOK, r.Code)
	assert.Equal(suite.T(), 1, len(actualMenu))
	assert.Equal(suite.T(), dummyMenu[0].MenuName, actualMenu[0].MenuName)
	assert.Nil(suite.T(), err)
}

func (suite *MenuControllerTestSuite) TestListMenuApi_Failed() {
	suite.listMenuUseCaseMock.On("ListMenu").Return(nil, errors.New("failed"))
	NewMenuController(
		suite.routerGroupMock,
		suite.createMenuUseCaseMock,
		suite.listMenuUseCaseMock,
		suite.deleteMenuUseCaseMock,
		suite.updateMenuUseCaseMock,
	)
	r := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/api/master/menu", nil)
	suite.routerMock.ServeHTTP(r, request)
	var actualResponse dto.ErrorCode
	response := r.Body.String()
	json.Unmarshal([]byte(response), &actualResponse)
	fmt.Println("ACTUAL RESPONSE", actualResponse)
	assert.Equal(suite.T(), http.StatusInternalServerError, r.Code)
	assert.Equal(suite.T(), "XX", actualResponse.Code)
}

func TestMenuControllerTestSuite(t *testing.T) {
	suite.Run(t, new(MenuControllerTestSuite))
}
