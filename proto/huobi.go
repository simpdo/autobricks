package proto

//HBPing 火币服务器ping包
type HBPing struct {
	Timestamp int64 `json:"ping"`
}

//HBPong ping消息回包
type HBPong struct {
	Timestamp int64 `json:"pong"`
}

//HBSubReq 订阅请求
type HBSubReq struct {
	Sub string `json:"sub"`
	Id  string `json:"id"`
}
