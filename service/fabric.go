package service

import (
	"gas-fabric-service/common/httputil"
	"gas-fabric-service/common/log"
	"gas-fabric-service/config"
	"gas-fabric-service/model"
)

var logger = log.GetLogger("service", log.ERROR)

type FabricService struct {
	Url string
}

func (f FabricService) QueryChaincode(chaincode model.FabricChaincode) []byte {
	return httputil.PostJson(f.Url+"/queryChaincode", chaincode)
}

func (f FabricService) InvokeChaincode(chaincode model.FabricChaincode) []byte {
	return httputil.PostJson(f.Url+"/invokeChaincode", chaincode)
}

func (f FabricService) TransactionGetStatus(chaincode model.FabricChaincode, txid string) []byte {
	return httputil.PostJson(f.Url+"/transaction/getStatusById?id="+txid, chaincode)
}

func NewFabricService() *FabricService {
	return &FabricService{
		Url: config.Config.GetString("BaasFabricEngine"),
	}
}
