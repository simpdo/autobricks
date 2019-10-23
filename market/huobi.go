package market

import (
	"fmt"
)

const (
	huobiId = "01079"
)

/////////ping包S->C////////////////

//HBPing 火币服务器ping包
type HBPing struct {
	Timestamp int64 `json:"ping"`
}

//HBPong ping消息回包
type HBPong struct {
	Timestamp int64 `json:"pong"`
}

///////////////订阅请求///////////////////

//HBSubReq 订阅请求
type HBSubReq struct {
	Sub string `json:"sub"`
	ID  string `json:"id"`
}

//NewHBSubReq create sub request
func NewHBSubReq(symbol string, step int) *HBSubReq {
	sub := fmt.Sprintf("market.%s.depth.step%d", symbol, step)
	return &HBSubReq{Sub: sub, ID: huobiId}
}

//HBSubResp 订阅回复
type HBSubResp struct {
	ID        string `json:"id"`
	Status    string `json:"status"`
	Subbed    string `json:"subbed"`
	Timestamp int64  `json:"ts"`
}

//HBQuotation 行情通知
type HBQuotation struct {
	Symbol    string `json:"ch"`
	Timestamp int64  `json:"ts"`
}

///////////////取消请求///////////////////
