package market

import (
	"fmt"
	"io"

	"golang.org/x/net/websocket"
)

//1. 协议解析
//2. 协议处理
type Proxy interface {
	Decode(r io.Reader) (map[string]interface{}, error)
	Handle(message map[string]interface{}) []byte

	Subscribe(symbol string, step int)
}

type MarketConfig struct {
	OriginURL string //http服务地址
	WssURL    string //wss服务地址
}

type Market struct {
	conn  *websocket.Conn
	proxy Proxy
}

func (m *Market) Start() {
	for {
		req, err := m.proxy.Decode(m.conn)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		m.proxy.Handle(req)
	}
}

func (m *Market) Subscribe(symbol string, step int) {
	m.proxy.EncodeSubscribe(symbol, step)
}
