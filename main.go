package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/cupen/auto-jkyz/actions"
	"github.com/cupen/auto-jkyz/config"
	"github.com/pelletier/go-toml"
)

const (
	URL_OF_Login = "https://hk.sz.gov.cn:8118/userPage/login"
)

var (
	devtoolsWsURL = flag.String("devtools-ws-url", "ws://localhost:9222/devtools/browser", "DevTools WebSsocket URL")
	fpath         = flag.String("config", "conf.toml", "configuation file")
)

func main() {
	parse_cli()
	conf := parse_conf(*fpath)
	allocatorContext, cancel := chromedp.NewRemoteAllocator(context.Background(), *devtoolsWsURL)
	defer cancel()

	root, cancel := chromedp.NewContext(allocatorContext)
	defer cancel()

	err := actions.Login(root, URL_OF_Login, conf.Account)
	panicIf(err)

	time.Sleep(60 * 10 * time.Second)
}

func parse_cli() {
	if *devtoolsWsURL == "" {
		panic(fmt.Errorf("empty devtools websocket url"))
	}
	if *fpath == "" {
		panic(fmt.Errorf("empty config file path"))
	}
	flag.Parse()
}

func parse_conf(fpath string) *config.Config {
	data, err := os.ReadFile(fpath)
	panicIf(err)
	conf := config.Config{}
	err = toml.Unmarshal(data, &conf)
	panicIf(err)
	return &conf
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
