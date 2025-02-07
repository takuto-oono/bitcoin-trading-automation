package usecase

import (
	"fmt"
	"net/http"
	"slices"

	"github.com/bitcoin-trading-automation/internal/bitflyer-api/api"
	apiModels "github.com/bitcoin-trading-automation/internal/bitflyer-api/api/models"
	"github.com/bitcoin-trading-automation/internal/config"
	"github.com/bitcoin-trading-automation/internal/models"
)

const (
	ProductCodeBTCJPY  = "BTC_JPY"
	ProductCodeXRPJPY  = "XRP_JPY"
	ProductCodeETHJPY  = "ETH_JPY"
	ProductCodeXLMJPY  = "XLM_JPY"
	ProductCOdeMONAJPY = "MONA_JPY"

	ProductCodeETHBTC   = "ETH_BTC"
	ProductCodeBCHBTC   = "BCH_BTC"
	ProductCodeFXBTCJPY = "FX_BTC_JPY"

	GetExecutionsDefaultCount = "100"

	ChildOrderTypeLimit  = "LIMIT"
	ChildOrderTypeMarket = "MARKET"

	SideBuy  = "BUY"
	SideSell = "SELL"

	TimeInForceGTC = "GTC"
	TimeInForceIOC = "IOC"
	TimeInForceFOK = "FOK"
)

type IBitflyerUseCase interface {
	// public API
	GetBoard(productCode string) (apiModels.Board, int, error)
	GetTicker(productCode string) (models.Ticker, int, error)
	GetExecutions(productCode, count, before, after string) ([]apiModels.Execution, int, error)
	GetBoardState(productCode string) (apiModels.BoardStatus, int, error)
	GetHealth(productCode string) (apiModels.Health, int, error)

	// private API
	GetBalance() ([]apiModels.Balance, int, error)
	GetCollateral() (apiModels.Collateral, int, error)
	PostSendChildOrder(productCode, ChildOrderType, side string, price int, size float64, MinuteToExpire int, TimeInForce string, isDry bool) (apiModels.ChildOrder, int, error)
	PostCancelChildOrder(productCode, ChildOrderID string, isDry bool) (int, error)
	GetChildOrders() ([]apiModels.ChildOrder, int, error)
}

type BitflyerUseCase struct {
	PublicAPI  api.PublicAPI
	PrivateAPI api.PrivateAPI
	Config     config.Config
}

func NewBitflyerUseCase(cfg config.Config) IBitflyerUseCase {
	return &BitflyerUseCase{
		PublicAPI:  api.NewPublicAPI(cfg),
		PrivateAPI: api.NewPrivateAPI(cfg),
		Config:     cfg,
	}
}

func (bu *BitflyerUseCase) GetBoard(productCode string) (apiModels.Board, int, error) {
	if productCode == "" {
		productCode = ProductCodeBTCJPY
	}
	if !validateProductCode(productCode) {
		return apiModels.Board{}, http.StatusBadRequest, fmt.Errorf("invalid product code: %s", productCode)
	}

	boards, err := bu.PublicAPI.GetBoard(productCode)
	if err != nil {
		return apiModels.Board{}, http.StatusInternalServerError, err
	}

	return boards, http.StatusOK, nil
}

func (bu *BitflyerUseCase) GetTicker(productCode string) (models.Ticker, int, error) {
	if productCode == "" {
		productCode = ProductCodeBTCJPY
	}
	if !validateProductCode(productCode) {
		return models.Ticker{}, http.StatusBadRequest, fmt.Errorf("invalid product code: %s", productCode)
	}

	t, err := bu.PublicAPI.GetTicker(productCode)
	if err != nil {
		return models.Ticker{}, http.StatusInternalServerError, err
	}

	timestamp, err := parseTimestamp(t.Timestamp)
	if err != nil {
		return models.Ticker{}, http.StatusInternalServerError, err
	}

	ticker := models.NewTicker(t.TickID, t.ProductCode, t.State, timestamp, t.BestBid, t.BestAsk, t.BestBidSize, t.BestAskSize, t.TotalBidDepth, t.TotalAskDepth, t.MarketBidSize, t.MarketAskSize, t.Ltp, t.Volume, t.VolumeByProduct)

	return ticker, http.StatusOK, nil
}

func (bu *BitflyerUseCase) GetExecutions(productCode, count, before, after string) ([]apiModels.Execution, int, error) {
	if productCode == "" {
		productCode = ProductCodeBTCJPY
	}
	if count == "" {
		count = GetExecutionsDefaultCount
	}

	if !validateProductCode(productCode) {
		return []apiModels.Execution{}, http.StatusBadRequest, fmt.Errorf("invalid product code: %s", productCode)
	}

	executions, err := bu.PublicAPI.GetExecutions(productCode, count, before, after)
	if err != nil {
		return []apiModels.Execution{}, http.StatusInternalServerError, err
	}

	return executions, http.StatusOK, nil
}

func (bu *BitflyerUseCase) GetBoardState(productCode string) (apiModels.BoardStatus, int, error) {
	if productCode == "" {
		productCode = ProductCodeBTCJPY
	}
	if !validateProductCode(productCode) {
		return apiModels.BoardStatus{}, http.StatusBadRequest, fmt.Errorf("invalid product code: %s", productCode)
	}

	boardState, err := bu.PublicAPI.GetBoardState(productCode)
	if err != nil {
		return apiModels.BoardStatus{}, http.StatusInternalServerError, err
	}

	return boardState, http.StatusOK, nil
}

func (bu *BitflyerUseCase) GetHealth(productCode string) (apiModels.Health, int, error) {
	if productCode == "" {
		productCode = ProductCodeBTCJPY
	}
	if !validateProductCode(productCode) {
		return apiModels.Health{}, http.StatusBadRequest, fmt.Errorf("invalid product code: %s", productCode)
	}

	health, err := bu.PublicAPI.GetHealth(productCode)
	if err != nil {
		return apiModels.Health{}, http.StatusInternalServerError, err
	}

	return health, http.StatusOK, nil
}

func (bu *BitflyerUseCase) GetBalance() ([]apiModels.Balance, int, error) {
	balance, err := bu.PrivateAPI.GetBalance()
	if err != nil {
		return []apiModels.Balance{}, http.StatusInternalServerError, err
	}

	return balance, http.StatusOK, nil
}

func (bu *BitflyerUseCase) GetCollateral() (apiModels.Collateral, int, error) {
	collateral, err := bu.PrivateAPI.GetCollateral()
	if err != nil {
		return apiModels.Collateral{}, http.StatusInternalServerError, err
	}

	return collateral, http.StatusOK, nil
}

func (bu *BitflyerUseCase) PostSendChildOrder(productCode, childOrderType, side string, price int, size float64, minuteToExpire int, timeInForce string, isDry bool) (apiModels.ChildOrder, int, error) {
	if productCode == "" {
		productCode = ProductCodeBTCJPY
	}
	if !validateProductCode(productCode) {
		return apiModels.ChildOrder{}, http.StatusBadRequest, fmt.Errorf("invalid product code: %s", productCode)
	}

	if childOrderType == "" {
		return apiModels.ChildOrder{}, http.StatusBadRequest, fmt.Errorf("child order type is required")
	}

	if childOrderType != ChildOrderTypeLimit && childOrderType != ChildOrderTypeMarket {
		return apiModels.ChildOrder{}, http.StatusBadRequest, fmt.Errorf("invalid child order type: %s", childOrderType)
	}

	if side == "" {
		return apiModels.ChildOrder{}, http.StatusBadRequest, fmt.Errorf("side is required")
	}

	if side != SideBuy && side != SideSell {
		return apiModels.ChildOrder{}, http.StatusBadRequest, fmt.Errorf("invalid side: %s", side)
	}

	if timeInForce == "" {
		return apiModels.ChildOrder{}, http.StatusBadRequest, fmt.Errorf("time in force is required")
	}

	if timeInForce != TimeInForceGTC && timeInForce != TimeInForceIOC && timeInForce != TimeInForceFOK {
		return apiModels.ChildOrder{}, http.StatusBadRequest, fmt.Errorf("invalid time in force: %s", timeInForce)
	}

	req := apiModels.SendChildOrderRequest{
		ProductCode:    productCode,
		ChildOrderType: childOrderType,
		Side:           side,
		Price:          price,
		Size:           size,
		MinuteToExpire: minuteToExpire,
		TimeInForce:    timeInForce,
	}

	childOrder, err := bu.PrivateAPI.PostSendChildOrder(req, isDry)
	if err != nil {
		return apiModels.ChildOrder{}, http.StatusInternalServerError, err
	}

	return childOrder, http.StatusOK, nil
}

func (bu *BitflyerUseCase) PostCancelChildOrder(productCode, childOrderID string, isDry bool) (int, error) {
	if productCode == "" {
		productCode = ProductCodeBTCJPY
	}
	if !validateProductCode(productCode) {
		return http.StatusBadRequest, fmt.Errorf("invalid product code: %s", productCode)
	}

	if childOrderID == "" {
		return http.StatusBadRequest, fmt.Errorf("child order id is required")
	}

	req := apiModels.CancelChildOrderRequest{
		ProductCode:  productCode,
		ChildOrderID: childOrderID,
	}

	if err := bu.PrivateAPI.PostCancelChildOrder(req, isDry); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (bu *BitflyerUseCase) GetChildOrders() ([]apiModels.ChildOrder, int, error) {
	childOrders, err := bu.PrivateAPI.GetChildOrders()
	if err != nil {
		return []apiModels.ChildOrder{}, http.StatusInternalServerError, err
	}

	return childOrders, http.StatusOK, nil
}

func validateProductCode(productCode string) bool {
	allowProductCodes := []string{
		ProductCodeBTCJPY,
		ProductCodeXRPJPY,
		ProductCodeETHJPY,
		ProductCodeXLMJPY,
		ProductCOdeMONAJPY,
		ProductCodeETHBTC,
		ProductCodeBCHBTC,
		ProductCodeFXBTCJPY,
	}

	return slices.Contains(allowProductCodes, productCode)
}
