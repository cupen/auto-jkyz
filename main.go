package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/cupen/auto-jkyz/actions"
	"github.com/cupen/auto-jkyz/config"
	"github.com/pelletier/go-toml"
)

const ()

var (
	devtoolsWsURL = flag.String("devtools-ws-url", "ws://localhost:9222/devtools/browser", "DevTools WebSsocket URL")
	fpath         = flag.String("config", "conf.toml", "configuation file")
)

func main() {
	parse_cli()
	conf := parse_conf(*fpath)
	start_chrome(conf)
	allocatorContext, cancel := chromedp.NewRemoteAllocator(context.Background(), *devtoolsWsURL)
	defer cancel()

	root, cancel := chromedp.NewContext(allocatorContext)
	defer cancel()

	err := actions.Login(root, conf.Account)
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

func start_chrome(cfg *config.Config) {
	profileDir := filepath.Join(os.TempDir(), "auto-jkyz")
	args := []string{
		"--remote-debugging-port=9222",
		"--user-data-dir=" + profileDir,
	}
	cmd := exec.Command(cfg.Chrome.GetPath(), args...)
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	log.Printf("started chrome with profile directory: %s", profileDir)
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
