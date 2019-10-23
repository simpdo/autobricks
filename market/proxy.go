package market

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
)

// Proxy 虚拟货币市场代理，主要功能
//1. 协议解析
//2. 协议处理
type Proxy interface {
	Decode(r io.Reader) (interface{}, error)
	Handle(message map[string]interface{}) []byte
}

//HuobiProxy 火币全球网
type HuobiProxy struct {
}

//Decode 火币协议解析
func (proxy *HuobiProxy) Decode(r io.Reader) (map[string]interface{}, error) {

	gzip, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}
	var resp map[string]interface{}

	decoder := json.NewDecoder(gzip)
	err = decoder.Decode(&resp)

	return resp, err
}

//Handle 消息处理函数
func (proxy *HuobiProxy) Handle(message map[string]interface{}) []byte {
	resp := map[string]interface{}{}

	if proxy.isQuotationPing(message) {
		resp["pong"] = message["ping"]
	} else if proxy.isPing(message) {
		resp["op"] = "pong"
		resp["ts"] = message["ts"]
	}

	data, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("HuobiProxy handle json marshal err: %s\n", err.Error())
		return nil
	}

	return data
}

//isQuotationPing 判断是否为行情ping包
func (proxy *HuobiProxy) isQuotationPing(message map[string]interface{}) bool {
	_, ok := message["ping"]
	return ok
}

//isPing 判断是否为资产及订单ping包
func (proxy *HuobiProxy) isPing(message map[string]interface{}) bool {
	val, ok := message["op"]
	if !ok {
		return false
	}

	op, ok := val.(string)
	if ok && op == "ping" {
		return true
	}

	return false
}

func (proxy *HuobiProxy) isSubResp(message map[string]interface{}) bool {
	_, ok := message["subbed"]
	return ok
}
