package proto

import (
	"fmt"
)

const (
	huobiId = "01079"
)

//HBPing 火币服务器ping包
type HBPing struct {
	Timestamp int64 `json:"ping"`
}

//HBPong ping消息回包
type HBPong struct {
	Timestamp int64 `json:"pong"`
}

func NewHBPong(ts int64) *HBPong {
	return &HBPong{Timestamp: ts}
}

//HBSubReq 订阅请求
type HBSubReq struct {
	Sub string `json:"sub"`
	Id  string `json:"id"`
}

func NewHBSubReq(symbol string, step int) *HBSubReq {
	sub := fmt.Sprintf("market.%s.depth.step%d", symbol, step)
	return &HBSubReq{Sub: sub, Id: huobiId}
}

//HBSubResp 订阅回复
type HBSubResp struct {
	Id        string `json:"id"`
	Status    string `json:"status"`
	Subbed    string `json:"subbed"`
	Timestamp int64  `json:"ts"`
}
