package main

import (
	"autobricks/net"
	"fmt"
)

const (
	//HuobiOriginURL 火币地址
	HuobiOriginURL = "http://api.huobi.pro"

	//HuobiWssURL 火币websocket地址
	HuobiWssURL = "wss://api.huobi.pro/ws"
)

//MarketConfig 市场配置
type MarketConfig struct {
	OriginURL string      //http服务地址
	WssURL    string      //wss服务地址
	decoder   net.Decoder //市场协议解析
}

//MarketProxy 市场代理，负责处理与市场之间的通信，和市场数据的管理
type MarketProxy struct {
	client *net.WSClient
}

//NewMarketProxy 生成市场代理
func NewMarketProxy(config *MarketConfig) *MarketProxy {
	client, err := net.NewWsClient(HuobiWssURL, HuobiOriginURL)
	if client == nil {
		fmt.Println(err.Error())
		return nil
	}

}

//Start 开始运行。保持和市场的通信
func (m *MarketProxy) Start() {
	for {

	}
}
