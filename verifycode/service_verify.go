package verifycode

import (
	"log"
	"time"

	"github.com/veryfi/veryfi-go/veryfi"
	"github.com/veryfi/veryfi-go/veryfi/scheme"
)

func GetV2(fpath string) (string, error) {
	timeout := 10 * time.Second
	client, err := veryfi.NewClientV7(&veryfi.Options{
		ClientID: "YOUR_CLIENT_ID",
		Username: "YOUR_USERNAME",
		APIKey:   "YOUR_API_KEY",
		HTTP: veryfi.HTTPOptions{
			Timeout: timeout,
			Retry: veryfi.RetryOptions{
				Count: 1,
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.ProcessDocumentURL(scheme.DocumentURLOptions{
		FileURL: fpath,
		DocumentSharedOptions: scheme.DocumentSharedOptions{
			Tags: []string{"electric", "repair", "ny"},
		},
	})
	if err != nil {
		return "", err
	}
	log.Printf("%+v", resp)
	log.Printf("%s", resp.OCRText)
	return resp.OCRText, nil
}
