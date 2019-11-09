package market

import (
	"autobricks/util"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"regexp"

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

type HuobiSubNotify struct {
	Channel   string   `json:"ch"`
	Timestamp int      `json:"ts"`
	Tick      TickData `json:"tick"`
}

//NewHuobiMarket create
func NewHuobiMarket() *Market {
	conn, err := websocket.Dial(HuobiWssURL, "", HuobiOriginURL)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &Market{conn: conn, proxy: HuobiProxy{ID: HuobiID}}
}

//HuobiProxy  struct
type HuobiProxy struct {
	ID string
}

//Decode decode net message
func (proxy HuobiProxy) Decode(r io.Reader, message interface{}) error {

	gzip, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}
	var resp map[string]interface{}

	decoder := json.NewDecoder(gzip)
	err = decoder.Decode(&resp)

	return resp, err
}

//Handle handle message
func (proxy HuobiProxy) Handle(message interface{}) []byte {
	resp := map[string]interface{}{}

	buff, err := json.Marshal(message)

	fmt.Println(string(buff))

	if proxy.isSubbedResp(resp) {
		fmt.Println("received subscribe response")
	} else if proxy.isSubbedNotify(resp) {
		fmt.Println("received a subscribe notify")
	} else if proxy.isPing(resp) {
		fmt.Println("received a subscribe notify")
	} else if proxy.isOrderPing(resp) {

	}

	return nil
}

func (proxy HuobiProxy) Subscribe(item string, step int) {
	var req = map[string]string{}
	req["id"] = proxy.ID
	req["sub"] = fmt.Sprintf("market.%s.depth.step%d", item, step)

	buff, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

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

func (proxy HuobiProxy) isSubbedResp(m map[string]interface{}) bool {
	_, ok := util.GetStrValue(m, "subbed")
	return ok
}

func (proxy HuobiProxy) isPing(m map[string]interface{}) bool {
	_, ok := util.GetStrValue(m, "ping")
	return ok
}

func (proxy HuobiProxy) isOrderPing(m map[string]interface{}) bool {
	val, _ := util.GetStrValue(m, "op")
	return val == "ping"
}
