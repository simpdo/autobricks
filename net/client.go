package net

import (
	"golang.org/x/net/websocket"
)

//WSClient websocket客户端
type WSClient struct {
	conn *websocket.Conn

	originURL string
	wsURL     string
}

//Connect   连接ws服务器
func (client *WSClient) Connect(url, origin string) error {
	conn, err := websocket.Dial(url, "", origin)
	if err != nil {
		return err
	}
	client.conn = conn

	//保存地址，断线线后重连
	client.originURL = origin
	client.wsURL = url

	return nil
}
