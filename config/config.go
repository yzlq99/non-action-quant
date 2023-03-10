package config

import (
	"sync"

	"github.com/yzlq99/eastmoneyapi/client"

	"github.com/spf13/viper"
)

var defaultConfigFile = "./configs/config.yaml"

type Config struct {
	EastMoneyClientConfig client.EastMoneyClientConfig
	BatTradeSpec          string
}

var conf *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		viper.SetConfigFile(defaultConfigFile)
		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}
		if err := viper.Unmarshal(&conf); err != nil {
			panic(err)
		}
	})
	return conf

}

func SetConfigPath(path string) {
	defaultConfigFile = path
}
