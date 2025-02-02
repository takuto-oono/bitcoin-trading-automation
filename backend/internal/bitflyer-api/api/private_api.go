package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/bitcoin-trading-automation/internal/api"
	"github.com/bitcoin-trading-automation/internal/bitflyer-api/api/models"
	"github.com/bitcoin-trading-automation/internal/config"
)

type PrivateAPI interface {
	GetBalance() ([]models.Balance, error)
	GetCollateral() (models.Collateral, error)

	PostSendChildOrder(req models.SendChildOrderRequest, isDry bool) (models.ChildOrder, error)
	PostCancelChildOrder(req models.CancelChildOrderRequest, isDry bool) error
	GetChildOrders() ([]models.ChildOrder, error)
}

func NewPrivateAPI(cfg config.Config) PrivateAPI {
	api := api.NewAPI(cfg)

	return &BitFlyerAPI{
		BaseUrl:   BaseUrl(cfg.BitFlyer.BaseEndPoint),
		ApiKey:    cfg.BitFlyer.ApiKey,
		ApiSecret: cfg.BitFlyer.ApiSecret,
		API:       *api,
	}
}

func (bitAPI *BitFlyerAPI) GetBalance() ([]models.Balance, error) {
	url, err := bitAPI.BaseUrl.GetBalanceUrl()
	if err != nil {
		return []models.Balance{}, err
	}

	headerMap, err := bitAPI.API.PrivateRequestHeader(StringTimeStamp(), http.MethodGet, url, nil)
	if err != nil {
		return []models.Balance{}, err
	}

	var resModel []models.Balance
	if err := bitAPI.API.Do(http.MethodGet, nil, &resModel, url, headerMap); err != nil {
		return []models.Balance{}, err
	}

	return resModel, nil
}

func (bitAPI *BitFlyerAPI) GetCollateral() (models.Collateral, error) {
	url, err := bitAPI.BaseUrl.GetCollateralUrl()
	if err != nil {
		return models.Collateral{}, err
	}

	headerMap, err := bitAPI.API.PrivateRequestHeader(StringTimeStamp(), http.MethodGet, url, nil)
	if err != nil {
		return models.Collateral{}, err
	}

	resModel := models.Collateral{}
	if err := bitAPI.API.Do(http.MethodGet, nil, &resModel, url, headerMap); err != nil {
		return models.Collateral{}, err
	}
	return resModel, nil
}

func (bitAPI *BitFlyerAPI) PostSendChildOrder(req models.SendChildOrderRequest, isDry bool) (models.ChildOrder, error) {
	url, err := bitAPI.BaseUrl.PostSendChildOrderUrl()
	if err != nil {
		return models.ChildOrder{}, err
	}

	headerMap, err := bitAPI.API.PrivateRequestHeader(StringTimeStamp(), http.MethodPost, url, nil)
	if err != nil {
		return models.ChildOrder{}, err
	}

	resModel := models.ChildOrder{}
	if !isDry {
		if err := bitAPI.API.Do(http.MethodPost, req, &resModel, url, headerMap); err != nil {
			return models.ChildOrder{}, err
		}
	}

	return resModel, nil
}

func (bitAPI *BitFlyerAPI) PostCancelChildOrder(req models.CancelChildOrderRequest, isDry bool) error {
	url, err := bitAPI.BaseUrl.PostCancelChildOrderUrl()
	if err != nil {
		return err
	}

	headerMap, err := bitAPI.API.PrivateRequestHeader(StringTimeStamp(), http.MethodPost, url, nil)
	if err != nil {
		return err
	}

	if !isDry {
		if err := bitAPI.API.Do(http.MethodPost, req, nil, url, headerMap); err != nil {
			return err
		}
	}

	return nil
}

func (bitAPI *BitFlyerAPI) GetChildOrders() ([]models.ChildOrder, error) {
	url, err := bitAPI.BaseUrl.GetChildOrdersUrl()
	if err != nil {
		return []models.ChildOrder{}, err
	}

	headerMap, err := bitAPI.API.PrivateRequestHeader(StringTimeStamp(), http.MethodGet, url, nil)
	if err != nil {
		return []models.ChildOrder{}, err
	}

	var resModel []models.ChildOrder
	if err := bitAPI.API.Do(http.MethodGet, nil, &resModel, url, headerMap); err != nil {
		return []models.ChildOrder{}, err
	}
	return resModel, nil
}

func StringTimeStamp() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}
