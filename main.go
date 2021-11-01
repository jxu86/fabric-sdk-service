package main

import (
	"fabric-sdk-service/common/gintool"
	"fabric-sdk-service/config"
	"fabric-sdk-service/controller"
	"fabric-sdk-service/service"
	"net/http"

	"github.com/gin-gonic/gin"
	// "strings"
)

func main() {
	fabricService := service.NewFabricService()
	chaincodeService := service.NewChaincodeService(fabricService)
	apiController := controller.NewApiController(
		chaincodeService,
		fabricService,
	)
	router := gin.New()
	// 允许使用跨域请求  全局中间件
	router.Use(Cors())
	router.Use(gintool.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api/v1")
	{
		api.GET("/test", apiController.Test)
		api.POST("/chaincode/query", apiController.ChaincodeQuery)
		api.POST("/chaincode/invoke", apiController.ChaincodeInvoke)
	}

	router.Run(":" + config.Config.GetString("ServicePort"))
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
