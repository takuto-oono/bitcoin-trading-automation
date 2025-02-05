package api

import (
	"net/http"

	"github.com/bitcoin-trading-automation/internal/models"
)

func (api *API) SlackNotificationHealthCheck() error {
	url, err := SlackNotification(api.Config.Url.SlackNotification).HealthCheck()
	if err != nil {
		return err
	}

	return api.Do(http.MethodGet, nil, nil, url, nil)
}

func (api *API) SlackNotificationPostMessage(reqBody models.SlackNotificationPostMessage) error {
	url, err := SlackNotification(api.Config.Url.SlackNotification).PostMessage()
	if err != nil {
		return err
	}

	return api.Do(http.MethodPost, reqBody, nil, url, nil)
}
