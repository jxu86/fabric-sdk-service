package client

import (
	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-protos-go/peer"
)

// 创建通道客户端请求参数
type SdkChlClient struct {
	ConfigFile  string
	ChannelID   string
	UserName    string
	Org         string
	ConfigBytes []byte
}

// 内部请求周转
type CcRequest struct {
	ChaincodeID string   // 链码ID
	Fcn         string   // 函数名称
	Args        [][]byte // 参数
	Peers       []string // 请求peers

	TransientMap map[string][]byte // 隐私数据
	//InvocationChain []*fab.ChaincodeCall // 暂时用不到
}

type CcResponse struct {
	TransactionID    string
	TxValidationCode peer.TxValidationCode
	ChaincodeStatus  int32
	Payload          []byte
}

type ClientBase struct {
	ConfigFile string
	ChannelID  string
	UserName   string
	Org        string
}

type SdkClient interface {
	Initialize([]byte) error
	Close() error
	Query(CcRequest) (CcResponse, error)
	Invoke(CcRequest) (CcResponse, error)
	QueryCCData(string, []string) (CcResponse, error)

	ChaincodeInfo([]string) (*peer.ChaincodeQueryResponse, error)
	LedgerInfo([]string) (*common.BlockchainInfo, error)
	BlockInfoByTx(string, []string) (*common.Block, error)
	TxInfoById(string, []string) (*peer.ProcessedTransaction, error)
}
