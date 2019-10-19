// myproject project main.go
package main

import (
	//"bytes"
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

}

func main() {
	ws, err := websocket.Dial(huobiWssURL, "", huobiOriginURL)
	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
		return
	}

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
	var ping map[string]interface{}
	decoder.Decode(&ping)

	fmt.Println(ping)

}
