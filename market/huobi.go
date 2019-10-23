package market

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"

	"golang.org/x/net/websocket"
)

const ()

const (
	//HuobiId 在火币服务器订阅时的id
	HuobiId = "01079"

	//HuobiOriginURL 火币地址
	HuobiOriginURL = "http://api.huobi.pro"

	//HuobiWssURL 火币websocket地址
	HuobiWssURL = "wss://api.huobi.pro/ws"
)

//NewMarketClient 生成市场代理
func NewHuobiMarket() *Market {
	conn, err := websocket.Dial(HuobiWssURL, "", HuobiOriginURL)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &Market{conn: conn, proxy: HuobiProxy{}}

}

//HuobiProxy 火币全球网
type HuobiProxy struct {
}

//Decode 火币协议解析
func (proxy HuobiProxy) Decode(r io.Reader) (map[string]interface{}, error) {

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
func (proxy HuobiProxy) Handle(m map[string]interface{}) []byte {
	resp := map[string]interface{}{}

	//订阅行情通知
	if _, ok := m["ch"]; ok {
		//通知处理，不需要回复服务器
		//
		return nil
	}

	//需要回复的服务器消息
	if _, ok := m["subbed"]; ok {
		//订阅行情回复
		val, _ := GetString(m, "status")
		if val == "ok" {
			fmt.Println("%v sub success!", m["subbed"])
		}
	} else if op, ok := GetString(m, "op"); ok && op == "ping" {
		//资产订单ping包
		resp["op"] = "pong"
		resp["ts"] = m["ts"]
	} else if _, ok := m["ping"]; ok {
		//行情平包
		resp["pong"] = m["ping"]
	}

	//回复服务器消息
	data, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("HuobiProxy handle json marshal err: %s\n", err.Error())
		return nil
	}

	return data
}

func GetString(m map[string]interface{}, key string) (string, bool) {
	_, ok := m[key]
	if !ok {
		return "", false
	}

	val, ok := m[key].(string)
	if !ok {
		return "", false
	}

	return val, true
}
