package client

import (
	"encoding/base64"
	"fmt"
	"github.com/pkg/errors"
	"path/filepath"
	"strings"

	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-protos-go/peer"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/core"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/test/metadata"
)

type Client struct {
	ClientBase

	sdk        *fabsdk.FabricSDK
	cliChannel *channel.Client
	cliLedger  *ledger.Client
	cliResMgmt *resmgmt.Client
}

func (cli *Client) Initialize(confBytes []byte) error {
	var provider core.ConfigProvider
	if cli.ConfigFile == "" {
		provider = config.FromRaw(confBytes, "yaml")
		//provider = config.FromReader(urlReader, "yaml")
	} else {
		configPath := filepath.Join(metadata.GetProjectPath(), cli.ConfigFile)
		// 通过配置文件初始化sdk
		provider = config.FromFile(configPath)
	}

	var err error
	cli.sdk, err = fabsdk.New(provider)
	if err != nil {
		return fmt.Errorf("failed to create SDK,Error=%s", err.Error())
	}

	// 创建通道客户端，用于query和invoke
	clientContext := cli.sdk.ChannelContext(cli.ChannelID, fabsdk.WithOrg(cli.Org), fabsdk.WithUser(cli.UserName))
	cli.cliChannel, err = channel.New(clientContext)
	if err != nil {
		return fmt.Errorf("Failed to create new channel cliChannel, Error=%s", err.Error())
	}
	//log.Println("Channel cliChannel created")

	cli.cliLedger, err = ledger.New(clientContext)
	if err != nil {
		return fmt.Errorf("Failed to create new ledger cliLedger, Error=%s", err.Error())
	}

	// Channel management client is responsible for managing channels (create/update)
	resmgmtContext := cli.sdk.Context(fabsdk.WithUser(cli.UserName), fabsdk.WithOrg(cli.Org))
	cli.cliResMgmt, err = resmgmt.New(resmgmtContext)
	if err != nil {
		return errors.WithMessage(err, "Failed to create new channel management client")
	}

	return nil
}

func (cli *Client) Close() error {
	cli.sdk.Close()
	return nil
}

// ********************************** Channel functions *****************************
func (cli *Client) Query(req CcRequest) (cResp CcResponse, err error) {
	r := channel.Request{}
	r.Fcn = req.Fcn
	r.Args = req.Args
	r.ChaincodeID = req.ChaincodeID
	r.TransientMap = req.TransientMap

	// Create a request (proposal) and send it
	resp, err := cli.cliChannel.Query(r, channel.WithTargetEndpoints(req.Peers...))
	if err != nil {
		return cResp, err
	}
	cResp.Payload = resp.Payload
	cResp.TransactionID = string(resp.TransactionID)
	cResp.ChaincodeStatus = resp.ChaincodeStatus
	cResp.TxValidationCode = resp.TxValidationCode

	return
}

func (cli *Client) Invoke(req CcRequest) (cResp CcResponse, err error) {
	r := channel.Request{}
	r.Fcn = req.Fcn
	r.Args = req.Args
	r.ChaincodeID = req.ChaincodeID

	resp, err := cli.cliChannel.Execute(r, channel.WithTargetEndpoints(req.Peers...))
	if err != nil {
		return cResp, err
	}

	cResp.Payload = resp.Payload
	cResp.TransactionID = string(resp.TransactionID)
	cResp.ChaincodeStatus = resp.ChaincodeStatus
	cResp.TxValidationCode = resp.TxValidationCode

	return
}

func (cli *Client) QueryCCData(ccid string, peers []string) (cResp CcResponse, err error) {
	args := [][]byte{[]byte(cli.ChannelID)}
	args = append(args, []byte(ccid))

	req := channel.Request{
		ChaincodeID: "lscc",
		Fcn:         "getccdata",
		//Fcn: "getcollectionsconfig",
		Args: args,
	}

	resp, err := cli.cliChannel.Query(req, channel.WithTargetEndpoints(peers...))
	if err != nil {
		return cResp, err
	}
	cResp.Payload = resp.Payload
	cResp.TransactionID = string(resp.TransactionID)
	cResp.ChaincodeStatus = resp.ChaincodeStatus
	cResp.TxValidationCode = resp.TxValidationCode

	return
}

// ********************************** Ledger functions *****************************
func (cli *Client) LedgerInfo(peers []string) (*common.BlockchainInfo, error) {
	info, err := cli.cliLedger.QueryInfo(ledger.WithTargetEndpoints(peers...))
	if err != nil {
		return nil, err
	}

	return info.BCI, nil
}

func (cli *Client) BlockInfoByTx(txId string, peers []string) (*common.Block, error) {
	fabTxid := fab.TransactionID(txId)
	return cli.cliLedger.QueryBlockByTxID(fabTxid, ledger.WithTargetEndpoints(peers...))
}

func (cli *Client) TxInfoById(txId string, peers []string) (*peer.ProcessedTransaction, error) {
	fabTxid := fab.TransactionID(txId)
	return cli.cliLedger.QueryTransaction(fabTxid, ledger.WithTargetEndpoints(peers...))
}

// ********************************** Resmgmt functions *****************************
func (cli *Client) ChaincodeInfo(peers []string) (*peer.ChaincodeQueryResponse, error) {

	queryInstallCC, err := cli.cliResMgmt.LifecycleQueryInstalledCC(resmgmt.WithTargetEndpoints(peers...),
		resmgmt.WithRetry(retry.DefaultResMgmtOpts))
	if err != nil {
		return nil, err
	}
	//fmt.Println(queryInstallCC)
	var ccQueryResp peer.ChaincodeQueryResponse
	for _, c := range queryInstallCC {
		for _, v := range c.References {
			//fmt.Println(v[0].Name, v[0].Version, strings.Split(c.PackageID, ":")[1])
			var ccInfo peer.ChaincodeInfo
			ccInfo.Name = v[0].Name
			ccInfo.Version = v[0].Version
			idBase64 := strings.Split(c.PackageID, ":")[1]
			id, err := base64.StdEncoding.DecodeString(idBase64)
			if err != nil {
				return nil, err
			}
			ccInfo.Id = id

			ccQueryResp.Chaincodes = append(ccQueryResp.Chaincodes, &ccInfo)
		}
	}

	return &ccQueryResp, nil
}
