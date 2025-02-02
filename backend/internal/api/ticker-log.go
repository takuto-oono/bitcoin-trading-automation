package api

import (
	"net/http"

	"github.com/bitcoin-trading-automation/internal/bitflyer-api/api/models"
)

func (api *API) TickerLogPostTicker(ticker models.Ticket) error {
	url, err := TickerLogServer(api.Config.Url.TickerLogServer).PostTickerLog()
	if err != nil {
		return err
	}

	if err := api.Do(http.MethodPost, ticker, nil, url, nil); err != nil {
		return err
	}

	return nil
}
