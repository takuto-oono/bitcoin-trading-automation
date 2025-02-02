package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/bitcoin-trading-automation/internal/config"
)

type API struct {
	Config config.Config
}

func NewAPI(cfg config.Config) *API {
	return &API{
		Config: cfg,
	}
}

// TODO: bitflyerのapiフォルダと実装がダブっているので、共通化できるか検討
func Do(method string, reqModel, resModel any, url string, headers map[string]string) error {
	reqJson, err := MarshalJson(reqModel)
	if err != nil {
		return err
	}

	res, err := request(method, url, reqJson, headers)
	if err != nil {
		return err
	}

	resJson, err := readResponse(res)
	if err != nil {
		return err
	}

	if resModel == nil {
		return nil
	}

	return json.Unmarshal(resJson, resModel)
}

func request(method, url string, body []byte, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	return http.DefaultClient.Do(req)
}

func readResponse(resp *http.Response) ([]byte, error) {
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println(err)
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		fmt.Printf("status code: %d, body: %s", resp.StatusCode, string(body))
		return nil, errors.New(string(body))
	}

	return body, nil
}

func MarshalJson(v any) ([]byte, error) {
	if v == nil || v == "" {
		return []byte{}, nil
	}
	return json.Marshal(v)
}
