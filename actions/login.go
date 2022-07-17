package actions

import (
	"context"
	"log"
	"os"

	"github.com/chromedp/chromedp"
	"github.com/cupen/auto-jkyz/config"
	"github.com/cupen/auto-jkyz/verifycode"
)

const (
	URL_OF_Login = "https://hk.sz.gov.cn:8118/userPage/login"
)

func Login(root context.Context, acc *config.Account) error {
	log := log.New(os.Stdout, "login", log.Default().Flags())
	url := URL_OF_Login
	printf := func(f string, args ...interface{}) func(context.Context) error {
		return func(context.Context) error {
			log.Printf(f, args...)
			return nil
		}
	}

	fpath_verifyCode := "tmp/verify_code.png"
	codeCh := catchVerifyCode(root, fpath_verifyCode)
	log.Printf("starting login")
	elemTips := "#winLoginNotice div.flex1 button"
	if err := chromedp.Run(root,
		chromedp.Navigate(url),
		chromedp.ActionFunc(printf("open site")),
		chromedp.WaitVisible(elemTips),
		chromedp.ActionFunc(sleep(1, 2)),
		chromedp.ActionFunc(printf("click tips")),
		chromedp.Click(elemTips),
		chromedp.ActionFunc(sleep(1, 2)),
		chromedp.WaitVisible("img#img_verify"),
		chromedp.ActionFunc(printf("input idtype")),
		chromedp.SetValue(`//select[@id="select_certificate"]`, acc.IDType, chromedp.BySearch),
		chromedp.ActionFunc(printf("input username")),
		chromedp.SendKeys("input#input_idCardNo", acc.Username),
		chromedp.ActionFunc(printf("input password")),
		chromedp.SendKeys("input#input_pwd", acc.Password),
	); err != nil {
		return err
	}

	// waiting verycode file
	<-codeCh
	code := verifycode.MustGet(fpath_verifyCode)
	if err := chromedp.Run(root,
		chromedp.ActionFunc(printf("input verifyCode")),
		chromedp.SendKeys("input#input_verifyCode", code),
	); err != nil {
		return err
	}
	return nil
}
