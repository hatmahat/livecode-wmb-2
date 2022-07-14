package main

import "go-api-with-gin2/delivery"

// import (
// 	"fmt"
// 	"go-api-with-gin2/config"
// 	"go-api-with-gin2/delivery"
// 	"go-api-with-gin2/delivery/middleware"
// 	"go-api-with-gin2/model"
// 	"go-api-with-gin2/utils"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

func main() {
	delivery.Server().Run()

	// routerEngine := gin.Default()
	// cfg := config.NewConfig()

	// tokenService := utils.NewTokenService(cfg.TokenConfig)

	// routerGroup := routerEngine.Group("/api")
	// routerGroup.POST("/auth/login", func(ctx *gin.Context) {
	// 	var user model.Credential
	// 	if err := ctx.BindJSON(&user); err != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{
	// 			"message": "can't bind struct",
	// 		})
	// 		return
	// 	}

	// 	if user.Username == "enigma" && user.Password == "12345" {
	// 		token, err := tokenService.CreateAccessToken(&user)
	// 		fmt.Println(user.Password)
	// 		if err != nil {
	// 			ctx.AbortWithStatus(http.StatusUnauthorized)
	// 			return
	// 		}
	// 		err = tokenService.StoreAccessToken(user.Username, token)
	// 		if err != nil {
	// 			ctx.AbortWithStatus(http.StatusUnauthorized)
	// 			return
	// 		}
	// 		ctx.JSON(http.StatusOK, gin.H{
	// 			"token": token,
	// 		})

	// 	} else {
	// 		ctx.AbortWithStatus(http.StatusUnauthorized)
	// 	}
	// })

	// protectedGroup := routerGroup.Group("master", middleware.NewTokenValidator(tokenService).RequireToken())

	// protectedGroup.GET("/customer", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"message": ctx.GetString("user-id"),
	// 	})
	// })

	// protectedGroup.GET("/product", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"message": ctx.GetString("user-id"),
	// 	})
	// })

	// err := routerEngine.Run("localhost:8888")
	// if err != nil {
	// 	panic(err)
	// }
}
