package handle

import (
	"encoding/json"
	"fmt"

	"golang.org/x/net/websocket"
)

func PingHandle(ws *websocket.Conn, ts interface{}) {

	var pong = map[string]interface{}{}
	pong["pong"] = ts

	buff, err := json.Marshal(&pong)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ws.Write(buff)
}
