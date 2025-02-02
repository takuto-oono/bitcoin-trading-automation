package api

import (
	"fmt"
	"net/url"
)

type BitFlyerAPI string
type TickerLogServer string

func (b BitFlyerAPI) GetTicker() (string, error) {
	u, err := url.Parse(fmt.Sprintf("%s/ticker", b))
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

func (t TickerLogServer) PostTickerLog() (string, error) {
	u, err := url.Parse(fmt.Sprintf("%s/ticker-logs", t))
	if err != nil {
		return "", err
	}
	return u.String(), nil
}
