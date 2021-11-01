package model

import (
	"strings"
)

type FabricChain struct {
	ChainName     string   `json:"chainName"`
	Account       string   `json:"account"`       //用户帐号
	Consensus     string   `json:"consensus"`     //共识
	PeersOrgs     []string `json:"peersOrgs"`     //参与组织 除了orderer
	OrderCount    int      `json:"orderCount"`    //orderer节点个数
	PeerCount     int      `json:"peerCount"`     //每个组织节点个数
	ChannelName   string   `json:"channelName"`   //channel 名
	TlsEnabled    string   `json:"tlsEnabled"`    //是否开启tls  true or false
	FabricVersion string   `json:"fabricVersion"` //fabric 版本
}

type FabricBlockChain struct {
	ChainName     string      `json:"chainName"`
	Account       string      `json:"account"`   //用户帐号
	Consensus     string      `json:"consensus"` //共识
	PeersOrgs     []FabricOrg `json:"peersOrgs"` //参与组织 除了orderer
	ExtraOrgs     []FabricOrg `json:"extraOrgs"`
	OrdererOrg    FabricOrg   `json:"ordererOrg"`
	ChannelName   string      `json:"channelName"`   //channel 名
	TlsEnabled    string      `json:"tlsEnabled"`    //是否开启tls  true or false
	FabricVersion string      `json:"fabricVersion"` //fabric 版本
}

type FabricOrg struct {
	Account       string `json:"account"`       // 账号名称
	OrgName       string `json:"orgName"`       // 组织名字
	OrgType       string `json:"orgType"`       // 组织类型，peer or orderer
	Count         int    `json:"count"`         // 节点数量
	Domain        string `json:"domain"`        // 域名
	EnableNodeOUs bool   `json:"enableNodeOUs"` //
	FabricVersion string `json:"fabricVersion"` //
}

func (f FabricChain) GetHostDomain(org string) string {
	return strings.ToLower(f.Account + f.ChainName + org)
}

type FabricChannel struct {
	FabricChain
	ChaincodeId    string   `json:"chaincodeId"`
	ChaincodePath  string   `json:"chaincodePath"`
	ChaincodeBytes []byte   `json:"chaincodeBytes"`
	Version        string   `json:"version"`
	Policy         string   `json:"policy"`
	Args           [][]byte `json:"args"`
}

type ArgJson struct {
	ChaincodeId string `json:"ccid"`
	Event       string `json:"event"`
	Args        string `json:"args"`
}

type FabricChaincode struct {
	FabricBlockChain
	ChaincodeId    string   `json:"chaincodeId"`
	ChaincodePath  string   `json:"chaincodePath"`
	GithubPath     string   `json:"githubPath"`
	ChaincodeBytes []byte   `json:"chaincodeBytes"`
	Version        string   `json:"version"`
	Policy         string   `json:"policy"`
	Args           [][]byte `json:"args"`
	ArgsStr        string   `json:"argsStr"`
	ArgJson        ArgJson  `json:"argJson"`
}

//ChaincodeCallRequest
type ChaincodeCallRequest struct {
	AreaCode     string      `json:"areaCode,omitempty"`
	ChaincodeId  string      `json:"ccid,omitempty"`
	FunctionName string      `json:"funcName,omitempty"`
	DocType      string      `json:"docType,omitempty"`
	Args         []string    `json:"args,omitempty"`
	Payload      interface{} `json:"payload" binding:"required"`
	Peers        []string    `json:"peers,omitempty"`
}

func (f FabricChannel) GetChain() FabricChain {
	return f.FabricChain
}
