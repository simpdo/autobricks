package market

import (
	"fmt"
	"io"

	"golang.org/x/net/websocket"
)

// Proxy 虚拟货币市场代理，主要功能
//1. 协议解析
//2. 协议处理
type Proxy interface {
	Decode(r io.Reader) (map[string]interface{}, error)
	Handle(message map[string]interface{}) []byte
}

//MarketConfig 市场配置
type MarketConfig struct {
	OriginURL string //http服务地址
	WssURL    string //wss服务地址
}

//MarketClient 市场代理，负责处理与市场之间的通信，和市场数据的管理
type Market struct {
	conn  *websocket.Conn
	proxy Proxy
}

//Start 开始运行。保持和市场的通信
func (m *Market) Start() {
	for {
		data, err := m.conn.Read()
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println(data)
	}
}
