package main

import (
	"gas-fabric-service/common/gintool"
	"gas-fabric-service/config"
	"gas-fabric-service/controller"
	"gas-fabric-service/service"
	"github.com/gin-gonic/gin"
	"net/http"
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

	api := router.Group("/api")
	{
		api.GET("/test", apiController.Test)

		api.POST("/gas/transaction/getTypes", apiController.GetTypes)
		api.GET("/gas/transaction/getTypes", apiController.GetTypes)
		api.POST("/gas/transaction/summary", apiController.TransactionSummary)
		api.POST("/gas/transaction/detail", apiController.TransactionDetail)

		api.POST("/chaincode/query", apiController.ChaincodeQuery)
		api.POST("/chaincode/invokegas", apiController.ChaincodeInvokeGas)
		api.POST("/chaincode/invokeBatch", apiController.ChaincodeInvokeBatchGas)

		api.GET("/gas/transaction/getStatus", apiController.TransactionStatus)
		api.POST("/gas/transaction/getStatusBatch", apiController.TransactionStatusBatch)
	}

	// dispatch := map[string]func(string) string{
	// 	"ttt": test,
	// }
	// dispatch["ttt"]("abc")

	router.Run(":" + config.Config.GetString("GasServicePort"))
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
