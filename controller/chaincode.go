package controller

import (
	"fabric-sdk-service/common/gintool"
	"fabric-sdk-service/common/util"
	"fabric-sdk-service/config"
	"fabric-sdk-service/model"

	"github.com/gin-gonic/gin"

	"os"
	"strings"
)

var (
	DefaultChaincodeId    = "gas"
	DefaultInvokeFuncName = "register"
	DefaultQueryFuncName  = "query"

	DefaultInvokePeers []string
	DefaultQueryPeers  []string
)

func init() {
	if peerStr := os.Getenv("DEFAULT_QUERY_PEERS"); peerStr != "" {
		DefaultQueryPeers = strings.Split(peerStr, ",")
	}

	if peerStr := os.Getenv("DEFAULT_INVOKE_PEERS"); peerStr != "" {
		DefaultInvokePeers = strings.Split(peerStr, ",")
	}

	if ccid := os.Getenv("DEFAULT_CHAINCODE_ID"); ccid != "" {
		DefaultChaincodeId = ccid
	}

	if funcName := os.Getenv("DEFAULT_INVOKE_FUNCNAME"); funcName != "" {
		DefaultInvokeFuncName = funcName
	}

	if funcName := os.Getenv("DEFAULT_QUERY_FUNCNAME"); funcName != "" {
		DefaultQueryFuncName = funcName
	}

	DefaultChaincodeId = config.Config.GetString("DefaultChaincode")

}

// @Summary 链码query接口
// @Description 链码query接口
// @Tags 链码
// @Accept json
// @Param ccid body string true "链码ID"
// @Param event body string true "接口名称"
// @Param args body string true "参数，用逗号隔开"
// @Success 200 {object} gintool.RespData desc
// @Router /api/v1/chaincode/query [post]
func (a *ApiController) ChaincodeQuery(ctx *gin.Context) {

	var argJson model.ArgJson
	if err := ctx.ShouldBindJSON(&argJson); err != nil {
		gintool.ResultFail(ctx, err)
		return
	}
	logger.Info("argJson:", argJson)

	var ccReq model.ChaincodeCallRequest

	if argJson.ChaincodeId != "" {
		ccReq.ChaincodeId = argJson.ChaincodeId
	} else {
		ccReq.ChaincodeId = DefaultChaincodeId
	}
	ccReq.Peers = DefaultInvokePeers

	ccReq.FunctionName = util.FirstUpper(argJson.Event)
	ccReq.Args = argJson.Args
	// ccReq.Args = append(ccReq.Args, argJson.Args)

	logger.Info("ccReq:", ccReq)
	resp, err := a.ChaincodeService.QueryChaincode(ccReq)
	if err != nil {
		gintool.ResultFail(ctx, "")
	} else {
		gintool.ResultOk(ctx, resp)
	}
}

// @Summary 链码invoke接口
// @Description 链码invoke接口
// @Tags 链码
// @Accept json
// @Param ccid body string true "链码ID"
// @Param event body string true "接口名称"
// @Param args body string true "参数，用逗号隔开"
// @Success 200 {object} gintool.RespData desc
// @Router /api/v1/chaincode/invoke [post]
func (a *ApiController) ChaincodeInvoke(ctx *gin.Context) {

	var argJson model.ArgJson
	if err := ctx.ShouldBindJSON(&argJson); err != nil {
		gintool.ResultFail(ctx, err.Error())
		return
	}
	logger.Info("argJson:", argJson)

	var ccReq model.ChaincodeCallRequest
	if argJson.ChaincodeId != "" {
		ccReq.ChaincodeId = argJson.ChaincodeId
	} else {
		ccReq.ChaincodeId = DefaultChaincodeId
	}

	ccReq.Peers = DefaultInvokePeers
	ccReq.FunctionName = util.FirstUpper(argJson.Event)
	ccReq.Args = argJson.Args
	// ccReq.Args = append(ccReq.Args, argJson.Args)

	logger.Info("ccReq:", ccReq)
	resp, err := a.ChaincodeService.InvokeChaincode(ccReq)
	if err != nil {
		gintool.ResultFail(ctx, err.Error())
	} else {
		gintool.ResultOk(ctx, resp)
	}
}
