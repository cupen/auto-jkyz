package actions

import (
	"context"
	"log"

	"github.com/chromedp/chromedp"
	"github.com/cupen/auto-jkyz/config"
	"github.com/cupen/auto-jkyz/verifycode"
)

func Login(root context.Context, url string, acc *config.Account) error {
	fpath_verifyCode := "tmp/verify_code.png"
	codeCh := catchVerifyCode(root, fpath_verifyCode)
	log.Printf("starting login")
	elemTips := "#winLoginNotice div.flex1 button"
	if err := chromedp.Run(root,
		chromedp.Navigate(url),
		chromedp.ActionFunc(printf("open login page")),
		chromedp.WaitVisible(elemTips),
		chromedp.ActionFunc(printf("click tips")),
		chromedp.ActionFunc(sleep(1, 2)),
		chromedp.Click(elemTips),
		chromedp.ActionFunc(printf("click tips")),
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
