package market

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/golang/glog"
)

const (
	HuobiUrl = "https://api.huobi.pro"
)

type HuobiHttpProxy struct {
	client *http.Client
}

func NewHuobiHttpProxy(proxyURL string) *HuobiHttpProxy {
	proxy, err := url.Parse(proxyURL)
	if err != nil {
		glog.Error(err.Error())
		return nil
	}

	httpClient := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns: 10,
			Proxy:        http.ProxyURL(proxy),
		},
		Timeout: time.Duration(10) * time.Second,
	}

	return &HuobiHttpProxy{client: httpClient}
}

func (proxy *HuobiHttpProxy) GetSymbols() {
	resp, err := proxy.client.Get("https://api.huobi.pro/v1/common/symbols")
	if err != nil {
		glog.Error(err.Error())
		return
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		glog.Error(err.Error())
		return
	}

	glog.Info(string(data))
}

func (proxy *HuobiHttpProxy) GetDepthData(step int) {
	reqUrl := fmt.Sprintf("https://api.huobi.pro/market/depth?symbol=btcusdt&type=step%d", step)
	resp, err := proxy.client.Get(reqUrl)
	if err != nil {
		glog.Error(err.Error())
		return
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		glog.Error(err.Error())
		return
	}

	glog.Info(string(data))
}
