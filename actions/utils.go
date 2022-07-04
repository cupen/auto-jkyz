package actions

import (
	"context"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

func sleep(min, max int) func(context.Context) error {
	return func(context.Context) error {
		secs := min + rand.Intn(max)
		log.Printf("sleep %d seconds", secs)
		time.Sleep(time.Duration(secs) * time.Second)
		return nil
	}
}

func printf(f string, args ...interface{}) func(context.Context) error {
	return func(context.Context) error {
		log.Printf(f, args...)
		return nil
	}
}

func catchVerifyCode(ctx context.Context, fpath string) chan []byte {
	downloadCh := make(chan bool)
	var requestId network.RequestID
	chromedp.ListenTarget(ctx, func(v interface{}) {
		switch ev := v.(type) {
		case *network.EventRequestWillBeSent:
			if strings.HasPrefix(ev.Request.URL, "https://hk.sz.gov.cn:8118/user/getVerify") {
				requestId = ev.RequestID
			}
		case *network.EventLoadingFinished:
			if requestId != "" && ev.RequestID == requestId {
				downloadCh <- true
				close(downloadCh)
				log.Printf("EventLoadingFinished: requestId:%v", ev.RequestID)
			}
		}
	})

	codeCh := make(chan []byte)
	go func() {
		<-downloadCh
		var buf []byte
		if err := chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			buf, err = network.GetResponseBody(requestId).Do(ctx)
			return err
		})); err != nil {
			log.Fatal(err)
		}

		if err := os.MkdirAll(filepath.Dir(fpath), 0o755); err != nil {
			log.Fatal(err)
		}
		if err := ioutil.WriteFile(fpath, buf, 0644); err != nil {
			log.Fatal(err)
		}
		codeCh <- buf
	}()
	return codeCh
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
