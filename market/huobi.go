package market

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"regexp"

	"adidos.cn/autobricks/util"
	"github.com/golang/glog"

	"golang.org/x/net/websocket"
)

const (
	//HuobiID id
	HuobiID = "01079"

	//HuobiOriginURL http url
	HuobiOriginURL = "http://api.huobi.pro"

	//HuobiWssURL wss url
	HuobiWssURL = "wss://api.huobi.pro/ws"
)

const (
	pingReq    = "ping"
	opField    = "op"
	subbedResp = "subbed"
)

//HuobiSubNotify 订阅通知
type HuobiSubNotify struct {
	Channel   string   `json:"ch"`
	Timestamp int      `json:"ts"`
	Tick      TickData `json:"tick"`
}

//HuobiSubscribeResp 订阅回复
type HuobiSubscribeResp struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Subbed string `json:"subbed"`
	Ts     int    `json:"ts"`
}

type HuobiProxy struct {
	conn *websocket.Conn
	ID   string
}

//NewHuobiMarket create
func NewHuobiProxy() (*HuobiProxy, error) {
	conn, err := websocket.Dial(HuobiWssURL, "", HuobiOriginURL)
	if err != nil {
		return nil, err
	}

	return &HuobiProxy{conn: conn, ID: HuobiID}, nil
}

//Subscribe 订阅请求
func (proxy *HuobiProxy) Subscribe(item string, step int) error {
	req := make(map[string]string)
	req["id"] = proxy.ID
	req["sub"] = fmt.Sprintf("market.%s.depth.step%d", item, step)

	buff, err := json.Marshal(req)
	if err != nil {
		return err
	}

	_, err = proxy.conn.Write(buff)
	if err != nil {
		return err
	}

	var resp HuobiSubscribeResp
	err = proxy.decode(&resp)
	if err != nil {
		return err
	}

	glog.Infof("subscribe %s response %v", item, resp)
	return nil
}

//Handle handle message
func (proxy *HuobiProxy) Start() {
	for {
		message := make(map[string]interface{})
		err := proxy.decode(message)
		if err != nil {
			glog.Errorf("huobi decode buffer error %s", err.Error())
			continue
		}

		buff, err := json.Marshal(message)
		glog.Infof("huobi message %s", string(buff))

		if proxy.isSubbedNotify(message) {
			glog.Info("received a subscribe notify")
		} else if proxy.isPing(message) {
			glog.Info("received a subscribe notify")
		} else if proxy.isOrderPing(message) {
			glog.Info("received a order ping")
		}
	}

	return
}

//Decode decode net message
func (proxy *HuobiProxy) decode(message interface{}) error {
	gzip, err := gzip.NewReader(proxy.conn)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(gzip)
	err = decoder.Decode(message)

	return err
}

func (proxy HuobiProxy) isSubbedNotify(m map[string]interface{}) bool {
	val, ok := util.GetStrValue(m, "ch")
	if !ok {
		return false
	}

	result, err := regexp.MatchString("market.*depth.*", val)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	return result
}

func (proxy HuobiProxy) isPing(m map[string]interface{}) bool {
	_, ok := util.GetStrValue(m, "ping")
	return ok
}

func (proxy HuobiProxy) isOrderPing(m map[string]interface{}) bool {
	val, _ := util.GetStrValue(m, "op")
	return val == "ping"
}
