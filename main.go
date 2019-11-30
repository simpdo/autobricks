// myproject project main.go
package main

import (
	"flag"
	"os"

	"adidos.cn/autobricks/market"
	"github.com/golang/glog"
)

func main() {
	flag.Parse()

	defer glog.Flush()
	glog.Infof("start success, pid %d", os.Getpid())

	huobi := market.NewHuobiHttpProxy("http://127.0.0.1:1081")
	if huobi == nil {
		return
	}

	huobi.GetSymbols()

	for i := 0; i < 6; i++ {
		huobi.GetDepthData(i)
	}
}
