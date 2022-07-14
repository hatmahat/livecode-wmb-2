package delivery

import (
	"fmt"
	"go-api-with-gin2/config"
	"go-api-with-gin2/delivery/controller"
	"go-api-with-gin2/delivery/middleware"
	"go-api-with-gin2/manager"
	"go-api-with-gin2/model"
	"go-api-with-gin2/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type appServer struct {
	useCaseManager manager.UseCaseManager
	engine         *gin.Engine
	host           string
	routerGroup    *gin.RouterGroup
}

func Server() *appServer {
	r := gin.Default()
	appConfig := config.NewConfig()
	infra := manager.NewInfra(&appConfig)
	repoManager := manager.NewRepoManager(infra)
	useCaseManager := manager.NewUseCaseManager(repoManager)
	host := appConfig.ApiConfig.Url

	tokenService := utils.NewTokenService(appConfig.TokenConfig)

	routerGroup := r.Group("/api")
	routerGroup.POST("/auth/login", func(ctx *gin.Context) {
		var user model.Credential
		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "can't bind struct",
			})
			return
		}

		if user.Username == "enigma" && user.Password == "12345" {
			token, err := tokenService.CreateAccessToken(&user)
			fmt.Println(user.Password)
			if err != nil {
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			err = tokenService.StoreAccessToken(user.Username, token)
			if err != nil {
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})

		} else {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
	})

	protectedGroup := routerGroup.Group("master", middleware.NewTokenValidator(tokenService).RequireToken())

	protectedGroup.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": ctx.GetString("user-id"),
		})
	})

	return &appServer{
		useCaseManager: useCaseManager,
		engine:         r,
		host:           host,
		routerGroup:    protectedGroup,
	}
}

func (a *appServer) initControllers() {
	controller.NewMenuController(a.routerGroup, a.useCaseManager.CreateMenuUseCase(), a.useCaseManager.ListMenuUseCase(), a.useCaseManager.DeleteMenuUseCase(), a.useCaseManager.UpdateMenuUseCase())
	controller.NewMenuPriceController(a.routerGroup, a.useCaseManager.CreateMenuPriceUseCase(), a.useCaseManager.ListMenuPriceUseCase(), a.useCaseManager.DeleteMenuPriceUseCase(), a.useCaseManager.UpdateMenuPriceUseCase())
	controller.NewBillDetailController(a.routerGroup, a.useCaseManager.CreateBillDetailUseCase(), a.useCaseManager.ListBillDetailUseCase())
	controller.NewBillController(a.routerGroup, a.useCaseManager.CreateBillUseCase(), a.useCaseManager.ListBillUseCase())
	controller.NewTableController(a.routerGroup, a.useCaseManager.CreateTableUseCase(), a.useCaseManager.ListTableUseCase(), a.useCaseManager.DeleteTableUseCase(), a.useCaseManager.UpdateTableUseCase())
	controller.NewTransTypeController(a.routerGroup, a.useCaseManager.CreateTransTypeUseCase(), a.useCaseManager.ListTransTypeUseCase(), a.useCaseManager.DeleteTransTypeUseCase(), a.useCaseManager.UpdateTransTypeUseCase())
	controller.NewCustomerController(a.routerGroup, a.useCaseManager.CreateCustomerUseCase(), a.useCaseManager.ListCustomerUseCase(), a.useCaseManager.UpdateCustomerUseCase())
	controller.NewDiscountController(a.routerGroup, a.useCaseManager.CreateDiscountUseCase(), a.useCaseManager.ListDiscountUseCase(), a.useCaseManager.DeleteDiscountUseCase(), a.useCaseManager.UpdateDiscountUseCase())
}

func (a *appServer) Run() {
	a.initControllers()
	err := a.engine.Run(a.host)
	if err != nil {
		panic(err)
	}
	balance, err := a.useCaseManager.CheckBalanceUseCaseLopei().GetBalance(int32(3))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("BALANCE", balance)
}
