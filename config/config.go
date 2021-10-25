package config

import (
	"gas-fabric-service/common/log"
	"os"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Config *viper.Viper

var logger = log.GetLogger("gas.config", log.INFO)

func init() {
	go watchConfig()
	loadConfig()
	logger.Infof("Env: %v", Config.AllSettings())
}

//监听配置改变
func watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		logger.Info("Config file changed:", e.Name)
		//改变重新加载
		loadConfig()
	})
}

//加载配置
func loadConfig() {
	viper.SetConfigName("gasconfig") // name of kubeconfig file
	viper.AddConfigPath(".")         // optionally look for kubeconfig in the working directory
	viper.AddConfigPath("/etc/baas") // path to look for the kubeconfig file in
	err := viper.ReadInConfig()      // Find and read the feconfig.yaml file
	if err != nil {                  // Handle errors reading the kubeconfig file
		logger.Errorf("Fatal error config file: %s \n", err)
		os.Exit(-1)
	}
	//全局配置
	Config = viper.GetViper()
	loadEnv()
	Config.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	Config.SetEnvKeyReplacer(replacer)
}

func loadEnv() {
	replaces := []string{
		"GasServicePort",
		"BaasFabricEngine",
		"BaasChannelArtifactsPath",
		"ConfigFileName",
		"ChannelId",
		"chainName",
	}

	for _, env := range replaces {
		if v := os.Getenv(env); v != "" {
			Config.Set(env, v)
		}
	}
}
