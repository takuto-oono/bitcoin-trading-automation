package api

import (
	"net/http"

	"github.com/bitcoin-trading-automation/internal/api"
	"github.com/bitcoin-trading-automation/internal/bitflyer-api/api/models"
	"github.com/bitcoin-trading-automation/internal/config"
)

type PublicAPI interface {
	GetBoard(productCode string) (models.Board, error)
	GetTicker(productCode string) (models.Ticker, error)
	GetExecutions(productCode string, count, before, after string) ([]models.Execution, error)
	GetBoardState(productCode string) (models.BoardStatus, error)
	GetHealth(productCode string) (models.Health, error)
}

// TODO アドレスを返すようにする
func NewPublicAPI(cfg config.Config) PublicAPI {
	api := api.NewAPI(cfg)

	return BitFlyerAPI{
		BaseUrl:   BaseUrl(cfg.BitFlyer.BaseEndPoint),
		ApiKey:    cfg.BitFlyer.ApiKey,
		ApiSecret: cfg.BitFlyer.ApiSecret,
		API:       *api,
	}
}

// TODO レシーバーをアドレスにする
func (bitAPI BitFlyerAPI) GetBoard(productCode string) (models.Board, error) {
	url, err := bitAPI.BaseUrl.GetBoardUrl(productCode)
	if err != nil {
		return models.Board{}, err
	}
	resModel := models.Board{}
	if err := bitAPI.API.Do(http.MethodGet, nil, &resModel, url, nil); err != nil {
		return models.Board{}, err

	}
	return resModel, nil
}

func (bitAPI BitFlyerAPI) GetTicker(productCode string) (models.Ticker, error) {
	url, err := bitAPI.BaseUrl.GetTickerUrl(productCode)
	if err != nil {
		return models.Ticker{}, err
	}
	resModel := models.Ticker{}
	if err := bitAPI.API.Do(http.MethodGet, nil, &resModel, url, nil); err != nil {
		return models.Ticker{}, err
	}
	return resModel, nil
}

func (bitAPI BitFlyerAPI) GetExecutions(productCode, count, before, after string) ([]models.Execution, error) {
	url, err := bitAPI.BaseUrl.GetExecutionsUrl(productCode, count, before, after)
	if err != nil {
		return []models.Execution{}, err
	}
	var resModel []models.Execution
	if err := bitAPI.API.Do(http.MethodGet, nil, &resModel, url, nil); err != nil {
		return []models.Execution{}, err
	}
	return resModel, nil
}

func (bitAPI BitFlyerAPI) GetBoardState(productCode string) (models.BoardStatus, error) {
	url, err := bitAPI.BaseUrl.GetBoardStateUrl(productCode)
	if err != nil {
		return models.BoardStatus{}, err
	}
	resModel := models.BoardStatus{}
	if err := bitAPI.API.Do(http.MethodGet, nil, &resModel, url, nil); err != nil {
		return models.BoardStatus{}, err
	}
	return resModel, nil
}

func (bitAPI BitFlyerAPI) GetHealth(productCode string) (models.Health, error) {
	url, err := bitAPI.BaseUrl.GetHealthUrl(productCode)
	if err != nil {
		return models.Health{}, err
	}
	resModel := models.Health{}
	if err := bitAPI.API.Do(http.MethodGet, nil, &resModel, url, nil); err != nil {
		return models.Health{}, err
	}
	return resModel, nil
}
