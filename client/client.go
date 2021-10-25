package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"gas-fabric-service/config"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	ChannelId = "c1"
	chainName = "gaschain"

	cliNum    = 1
	clients   []SdkClient
	cliConfig []SdkChlClient

	USERNAME       = "Admin"
	configFileName = "c1.yaml"
	defaultOrgs    = []string{"p1"}
	orgNum         = 0

	urlConfigPath = ""
	configMode    = ""
)

func init() {
	fmt.Println("init start ...")
	if numStr := os.Getenv("MAX_CLIENT_NUM"); numStr != "" {
		if n, err := strconv.Atoi(numStr); err == nil {
			cliNum = n
		}
	}

	if chlStr := os.Getenv("DEFAULT_CHANNEL_ID"); chlStr != "" {
		ChannelId = chlStr
	}
	if chStr := os.Getenv("DEFAULT_CHAIN_NAME"); chStr != "" {
		chainName = chStr
	}
	if nameStr := os.Getenv("DEFAULT_CONFIG_NAME"); nameStr != "" {
		configFileName = nameStr
	}
	if urlStr := os.Getenv("URL_CONFIG_PATH"); urlStr != "" {
		urlConfigPath = urlStr
	}
	if modStr := os.Getenv("DEFAULT_CONFIG_MODE"); modStr != "" {
		configMode = modStr
	}

	if numStr := os.Getenv("DEFAULT_ORG"); numStr != "" {
		defaultOrgs = strings.Split(numStr, ",")
	}

	configFileName = config.Config.GetString("ConfigFileName")
	ChannelId = config.Config.GetString("ChannelId")
	chainName = config.Config.GetString("ChainName")

	cfgTemplate := SdkChlClient{}
	if configMode == "URL" {
		cfgTemplate.ChannelID = ChannelId
		cfgTemplate.UserName = USERNAME
		cfgByte, err := getConfigByUrl(urlConfigPath)
		if err != nil {
			fmt.Println("getConfigByUrl", err)
			return
		}
		cfgTemplate.ConfigBytes = cfgByte
	} else {
		// cfgTemplate.ConfigFile = "./channel-artifacts/" + configFileName

		cfgTemplate.ConfigFile = config.Config.GetString("BaasChannelArtifactsPath") + configFileName
		cfgTemplate.ChannelID = ChannelId
		cfgTemplate.UserName = USERNAME
	}
	fmt.Println("ConfigFile: ", cfgTemplate.ConfigFile)
	orgNum = len(defaultOrgs)
	for i := 0; i < orgNum; i++ {
		cfgTemplate.Org = defaultOrgs[i%orgNum]
		cliConfig = append(cliConfig, cfgTemplate)
	}

	generate()
}

func generate() {
	fmt.Println("generate start...")
	for i := 0; i < cliNum*orgNum; i++ {
		if cli, err := clientFactory(i); err == nil {
			clients = append(clients, cli)
		}
	}
}

func NewClient(ccc SdkChlClient) (SdkClient, error) {

	fmt.Println("gm client...")
	client := gmClient{}
	/// 赋值
	client.ConfigFile = ccc.ConfigFile
	client.ChannelID = ccc.ChannelID
	client.UserName = ccc.UserName
	client.Org = ccc.Org

	err := client.Initialize(ccc.ConfigBytes)
	if err != nil {
		return nil, err
	}

	return &client, nil

	// if os.Getenv("FABRIC_SDK_MODE") == "gm" {
	// 	fmt.Println("gm client...")
	// 	client := gmClient{}
	// 	/// 赋值
	// 	client.ConfigFile = ccc.ConfigFile
	// 	client.ChannelID = ccc.ChannelID
	// 	client.UserName = ccc.UserName
	// 	client.Org = ccc.Org

	// 	err := client.Initialize(ccc.ConfigBytes)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	return &client, nil
	// } else {
	// 	fmt.Println("normal client...")
	// 	client := Client{}
	// 	/// 赋值
	// 	client.ConfigFile = ccc.ConfigFile
	// 	client.ChannelID = ccc.ChannelID
	// 	client.UserName = ccc.UserName
	// 	client.Org = ccc.Org

	// 	err := client.Initialize(ccc.ConfigBytes)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	return &client, nil
	// }
}

func clientFactory(idx int) (SdkClient, error) {
	cli, err := NewClient(cliConfig[idx%orgNum])
	if err != nil {
		fmt.Println("Failed to create new SDK: ", err)
		return nil, err
	}
	return cli, nil
}

func GetClientRand() (SdkClient, error) {
	if len(clients) == 0 {
		generate()
		return nil, errors.New("BC client connect fail, try it later")
	}
	randIdx := 0
	if cliNum > 1 {
		randIdx = rand.Intn(cliNum)
	}
	return clients[randIdx], nil
}

func GetChannelId() string {
	return ChannelId
}

type UrlConfigRequest struct {
	Orgs        string `json:"orgs,omitempty"`
	ChannelName string `json:"channelName,omitempty"`
	Name        string `json:"name,omitempty"`
}

func getConfigByUrl(url string) (out []byte, err error) {
	confReq := UrlConfigRequest{
		Orgs:        strings.Join(defaultOrgs, ","),
		ChannelName: ChannelId,
		Name:        chainName,
	}

	confReqByte, err := json.Marshal(confReq)
	if err != nil {
		return nil, err
	}

	payload := strings.NewReader(string(confReqByte))
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, err
	}
	//req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 3 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.ContentLength < 1 {
		return nil, errors.New("can't fetch valid response")
	}

	out, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return
}
