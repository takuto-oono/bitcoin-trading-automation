package api

import "github.com/bitcoin-trading-automation/internal/api"

type BaseUrl string

type BitFlyerAPI struct {
	BaseUrl   BaseUrl
	ApiKey    string
	ApiSecret string
	API       api.API
}
