package service

import (
	"gas-fabric-service/client"
	"gas-fabric-service/model"
)

type ChaincodeService struct {
	FabircService *FabricService
}

func (l *ChaincodeService) QueryChaincode(ccReq model.ChaincodeCallRequest) (interface{}, error) {
	cli, err := client.GetClientRand()
	if err != nil {
		return nil, err
	}

	var args [][]byte
	for _, a := range ccReq.Args {
		args = append(args, []byte(a))
	}

	resp, err := cli.Query(client.CcRequest{
		ChaincodeID: ccReq.ChaincodeId,
		Fcn:         ccReq.FunctionName,
		Args:        args,
		Peers:       ccReq.Peers,
	})
	if err != nil {
		return nil, err
	}
	return string(resp.Payload), nil
}

func (l *ChaincodeService) InvokeChaincode(ccReq model.ChaincodeCallRequest) (interface{}, error) {

	cli, err := client.GetClientRand()
	if err != nil {
		return nil, err
	}

	var byteArgs [][]byte
	for _, arg := range ccReq.Args {
		byteArgs = append(byteArgs, []byte(arg))
	}

	// logger.Info("byteArgs: ", byteArgs)
	resp, err := cli.Invoke(client.CcRequest{
		ChaincodeID: ccReq.ChaincodeId,
		Fcn:         ccReq.FunctionName,
		Args:        byteArgs,
		Peers:       ccReq.Peers,
	})
	if err != nil {
		return nil, err
	}

	// res := map[string]interface{}{
	// 	"txId": string(resp.TransactionID),
	// }

	return string(resp.TransactionID), nil
}

func (l *ChaincodeService) QueryTransactionStatusByTxId(txid string, peers []string) error {

	cli, err := client.GetClientRand()
	if err != nil {
		return err
	}
	_, err = cli.TxInfoById(txid, peers)
	if err != nil {
		return err
	}
	return nil

}

func NewChaincodeService(fabircService *FabricService) *ChaincodeService {
	return &ChaincodeService{
		FabircService: fabircService,
	}
}
