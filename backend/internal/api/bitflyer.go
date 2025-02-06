package api

import (
	"net/http"

	"github.com/bitcoin-trading-automation/internal/bitflyer-api/api/models"
)

func (api *API) BitFlyerAPIHealthCheck() error {
	url, err := BitFlyerAPI(api.Config.Url.BitFlyerAPI).HealthCheck()
	if err != nil {
		return err
	}

	return api.Do(http.MethodGet, nil, nil, url, nil)
}

func (api *API) GetTicker() (models.Ticker, error) {
	url, err := BitFlyerAPI(api.Config.Url.BitFlyerAPI).GetTicker()
	if err != nil {
		return models.Ticker{}, err
	}

	var ticker models.Ticker
	if err := api.Do(http.MethodGet, nil, &ticker, url, nil); err != nil {
		return models.Ticker{}, err
	}

	return ticker, nil
}
