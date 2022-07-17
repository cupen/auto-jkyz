package actions

import (
	"context"
	"log"
	"os"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
)

const (
	URL_OF_ORDER = "https://hk.sz.gov.cn:8118/passInfo/detail"
)

func MakeOrder(root context.Context) error {
	log := log.New(os.Stdout, "makeorder", log.Default().Flags())
	url := URL_OF_ORDER
	printf := func(f string, args ...interface{}) func(context.Context) error {
		return func(context.Context) error {
			log.Printf(f, args...)
			return nil
		}
	}
	log.Printf("starting make order")
	if err := chromedp.Run(root,
		chromedp.Navigate(url),
		chromedp.WaitReady("#divSzArea .orange button"),
		chromedp.ActionFunc(printf("click order button")),
		// chromedp.Query()
		chromedp.QueryAfter("#divSzArea .orange button", func(c context.Context, id runtime.ExecutionContextID, nodes ...*cdp.Node) error {
			for _, node := range nodes {
				log.Printf(" node = %+v", node)
			}
			return nil
		}),
	); err != nil {
		return err
	}
	return nil
}
