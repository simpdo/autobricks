// myproject project main.go
package main

import (
	"sync"

	//"bytes"
	"autobricks/handle"
	"compress/gzip"
	"encoding/json"
	"fmt"

	//"io/ioutil"

	"golang.org/x/net/websocket"
)

const (
	huobiOriginURL = "http://api.huobi.pro"
	huobiWssURL    = "wss://api.huobi.pro/ws"
)

func onMessage(ws *websocket.Conn) {
	for {
		frameReader, err := ws.NewFrameReader()
		if err != nil {
			fmt.Printf("err: %s\n", err.Error())
			return
		}

		frame, err := ws.HandleFrame(frameReader)
		if err != nil {
			fmt.Printf("err: %s\n", err.Error())
			return
		}

		gzip, err := gzip.NewReader(frame)
		if err != nil {
			fmt.Printf("err: %s\n", err.Error())
			return
		}
		decoder := json.NewDecoder(gzip)
		var resp map[string]interface{}
		decoder.Decode(&resp)
		if ts, ok := resp["ping"]; ok {
			handle.PingHandle(ws, ts)
		}

		fmt.Println(resp)
	}
}

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
	go onMessage(ws)

	wg.Wait()
}
