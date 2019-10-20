package main

import (
	"autobricks/net"
)

const (
	HuobiOriginURL = "http://api.huobi.pro"
	HuobiWssURL    = "wss://api.huobi.pro/ws"
)

type Market struct {
	client *net.WSClient
}

func (m *Market) Init() {
	m.client = net.NewWsClient(HuobiWssURL, HuobiOriginURL)
	if m.client == nil {
		return
	}

}

func (m *Market) Run() {
	for {
		m.client.r
	}
}
