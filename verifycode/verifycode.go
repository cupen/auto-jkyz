package verifycode

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type VerifyCodeProvider interface {
	GetVerifyCode([]byte) (string, error)
}

func requestPOST(url string, body []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func Get(fpath string) (string, error) {
	url := "http://localhost:8080/base64"
	data, err := os.ReadFile(fpath)
	if err != nil {
		return "", err
	}
	reqBody, _ := json.Marshal(map[string]interface{}{
		"base64": base64.StdEncoding.EncodeToString(data),
		"trim":   "\n",
	})
	resp, err := requestPOST(url, reqBody)
	if err != nil {
		return "", err
	}
	var _resp = struct {
		Version string
		Result  string
	}{}
	if err := json.Unmarshal(resp, &_resp); err != nil {
		return "", err
	}
	log.Printf("%s", string(resp))
	if _resp.Result == "" {
		return "", fmt.Errorf("invalid response")
	}
	return _resp.Result, nil
	// client := gosseract.NewClient()
	// defer client.Close()
	// client.SetImage(fpath)
	// return client.Text()
	return "", nil
}

func MustGet(fpath string) string {
	code, err := Get(fpath)
	if err != nil {
		panic(err)
	}
	return code
}
