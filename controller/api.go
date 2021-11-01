package controller

import (
	"fabric-sdk-service/common/gintool"
	"fabric-sdk-service/common/log"
	"fabric-sdk-service/service"
	"github.com/gin-gonic/gin"
)

var logger = log.GetLogger("controller", log.INFO)

// type ApiController struct {
// }

// func NewApiController() *ApiController {
// 	return &ApiController{}
// }

type ApiController struct {
	ChaincodeService *service.ChaincodeService
	FabricService    *service.FabricService
}

func NewApiController(
	chaincodeService *service.ChaincodeService,
	fabricService *service.FabricService,
) *ApiController {
	return &ApiController{
		ChaincodeService: chaincodeService,
		FabricService:    fabricService,
	}
}

func (a *ApiController) Test(ctx *gin.Context) {
	logger.Info("Test start ...")

	gintool.ResultOk(ctx, "test success!")
}
