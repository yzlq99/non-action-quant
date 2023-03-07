package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/yzlq99/eastmoneyapi/client"
	"github.com/yzlq99/non-action-quant/config"
)

var configPath string

func init() {
	log.SetFlags(log.Lshortfile)
	flag.StringVar(&configPath, "config", "", "")
	flag.Parse()
	if configPath != "" {
		config.SetConfigPath(configPath)
	}
}

func main() {

	c := client.NewEastMoneyClient(config.GetConfig().EastMoneyClientConfig)

	go func() {
		for {
			time.Sleep(time.Second * 2)
			res, err := c.GetStockList()
			if err != nil {
				panic(err)
			}
			str, _ := json.Marshal(res)
			fmt.Println(string(str))
		}
	}()
	select {}
}
