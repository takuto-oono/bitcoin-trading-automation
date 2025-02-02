package api

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
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

func (api *API) Do(method string, reqModel, resModel any, url string, headers map[string]string) error {
	reqJson, err := marshalJson(reqModel)
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

// https://lightning.bitflyer.com/docs#%E8%AA%8D%E8%A8%BC:~:text=%E4%BA%86%E6%89%BF%E3%81%8F%E3%81%A0%E3%81%95%E3%81%84%E3%80%82-,%E8%AA%8D%E8%A8%BC,-Private%20API%20%E3%81%AE
func (api *API) PrivateRequestHeader(timeStamp, method, url string, body []byte) (map[string]string, error) {
	path, err := extractPath(url)
	if err != nil {
		return nil, err
	}

	rawMessage := json.RawMessage(body)

	text := timeStamp + method + path + string(rawMessage)

	h := hmac.New(sha256.New, []byte(api.Config.BitFlyer.ApiSecret))
	h.Write([]byte(text))
	sign := hex.EncodeToString(h.Sum(nil))

	return map[string]string{
		"ACCESS-KEY":       api.Config.BitFlyer.ApiKey,
		"ACCESS-TIMESTAMP": timeStamp,
		"ACCESS-SIGN":      sign,
	}, nil
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

func marshalJson(v any) ([]byte, error) {
	if v == nil || v == "" {
		return []byte{}, nil
	}
	return json.Marshal(v)
}
