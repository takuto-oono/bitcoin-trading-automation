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

func (api *API) GetTicker() (models.Ticket, error) {
	url, err := BitFlyerAPI(api.Config.Url.BitFlyerAPI).GetTicker()
	if err != nil {
		return models.Ticket{}, err
	}

	var ticker models.Ticket
	if err := api.Do(http.MethodGet, nil, &ticker, url, nil); err != nil {
		return models.Ticket{}, err
	}

	return ticker, nil
}
