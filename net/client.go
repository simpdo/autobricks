package net

import (
	"io"

	"golang.org/x/net/websocket"
)

type Decoder interface {
	Decode(r io.Reader) (interface{}, error)
}

//WSClient websocket客户端
type WSClient struct {
	conn *websocket.Conn

	originURL string
	wsURL     string
}

type RespFunc func(data []byte)

//Connect   连接ws服务器
func NewWsClient(url, origin string) (*WSClient, error) {
	conn, err := websocket.Dial(url, "", origin)
	if err != nil {
		return nil, err
	}

	client := WSClient{}
	client.conn = conn

	//保存地址，断线线后重连
	client.originURL = origin
	client.wsURL = url

	return &client, nil
}

func (client *WSClient) Read(decoder Decoder) (interface{}, error) {
	frameReader, err := client.conn.NewFrameReader()
	if err != nil {
		return nil, err
	}

	frame, err := client.conn.HandleFrame(frameReader)
	if err != nil {
		return nil, err
	}

	resp, err := decoder.Decode(frame)

	return resp, nil
}
