// myproject project main.go
package main

import (
	"autobricks/net"
	"sync"

	//"bytes"
	"autobricks/handle"
	"encoding/json"
	"fmt"

	//"io/ioutil"

	"golang.org/x/net/websocket"
)

func Sub(ws *websocket.Conn) {
	var subReq = map[string]string{}
	subReq["sub"] = "market.btcusdt.depth.step0"
	subReq["id"] = "01079"

	data, _ := json.Marshal(subReq)

	ws.Write(data)
}

func main() {
	ws, err := websocket.Dial(huobiWssURL, "", huobiOriginURL)
	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)

	client := net.NewWsClient(huobiWssURL, HuobiOriginURL)
	if client == nil {
		fmt.Println("connect to server failed")
		wg.Done()
	}

	client.read()

	wg.Wait()
}
