package usecase

import (
	"fmt"
	"slices"

	"github.com/bitcoin-trading-automation/internal/bitflyer-api/api"
	"github.com/bitcoin-trading-automation/internal/bitflyer-api/api/models"
	"github.com/bitcoin-trading-automation/internal/config"
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
	GetBoard(productCode string) (models.Board, error)
	GetTicker(productCode string) (models.Ticket, error)
	GetExecutions(productCode, count, before, after string) ([]models.Execution, error)
	GetBoardState(productCode string) (models.BoardStatus, error)
	GetHealth(productCode string) (models.Health, error)

	// private API
	GetBalance() ([]models.Balance, error)
	GetCollateral() (models.Collateral, error)
	PostSendChildOrder(productCode, ChildOrderType, side string, price int, size float64, MinuteToExpire int, TimeInForce string, isDry bool) (models.ChildOrder, error)
	PostCancelChildOrder(productCode, ChildOrderID string, isDry bool) error
	GetChildOrders() ([]models.ChildOrder, error)
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

func (bu *BitflyerUseCase) GetBoard(productCode string) (models.Board, error) {
	if productCode == "" {
		productCode = ProductCodeBTCJPY
	}
	if !validateProductCode(productCode) {
		return models.Board{}, fmt.Errorf("invalid product code: %s", productCode)
	}
	return bu.PublicAPI.GetBoard(productCode)
}

func (bu *BitflyerUseCase) GetTicker(productCode string) (models.Ticket, error) {
	if productCode == "" {
		productCode = ProductCodeBTCJPY
	}
	if !validateProductCode(productCode) {
		return models.Ticket{}, fmt.Errorf("invalid product code: %s", productCode)
	}
	return bu.PublicAPI.GetTicker(productCode)
}

func (bu *BitflyerUseCase) GetExecutions(productCode, count, before, after string) ([]models.Execution, error) {
	if productCode == "" {
		productCode = ProductCodeBTCJPY
	}
	if count == "" {
		count = GetExecutionsDefaultCount
	}

	if !validateProductCode(productCode) {
		return []models.Execution{}, fmt.Errorf("invalid product code: %s", productCode)
	}

	return bu.PublicAPI.GetExecutions(productCode, count, before, after)
}

func (bu *BitflyerUseCase) GetBoardState(productCode string) (models.BoardStatus, error) {
	if productCode == "" {
		productCode = ProductCodeBTCJPY
	}
	if !validateProductCode(productCode) {
		return models.BoardStatus{}, fmt.Errorf("invalid product code: %s", productCode)
	}
	return bu.PublicAPI.GetBoardState(productCode)
}

func (bu *BitflyerUseCase) GetHealth(productCode string) (models.Health, error) {
	if productCode == "" {
		productCode = ProductCodeBTCJPY
	}
	if !validateProductCode(productCode) {
		return models.Health{}, fmt.Errorf("invalid product code: %s", productCode)
	}
	return bu.PublicAPI.GetHealth(productCode)
}

func (bu *BitflyerUseCase) GetBalance() ([]models.Balance, error) {
	return bu.PrivateAPI.GetBalance()
}

func (bu *BitflyerUseCase) GetCollateral() (models.Collateral, error) {
	return bu.PrivateAPI.GetCollateral()
}

func (bu *BitflyerUseCase) PostSendChildOrder(productCode, childOrderType, side string, price int, size float64, minuteToExpire int, timeInForce string, isDry bool) (models.ChildOrder, error) {
	if productCode == "" {
		productCode = ProductCodeBTCJPY
	}
	if !validateProductCode(productCode) {
		return models.ChildOrder{}, fmt.Errorf("invalid product code: %s", productCode)
	}

	if childOrderType == "" {
		return models.ChildOrder{}, fmt.Errorf("child order type is required")
	}

	if childOrderType != ChildOrderTypeLimit && childOrderType != ChildOrderTypeMarket {
		return models.ChildOrder{}, fmt.Errorf("invalid child order type: %s", childOrderType)
	}

	if side == "" {
		return models.ChildOrder{}, fmt.Errorf("side is required")
	}

	if side != SideBuy && side != SideSell {
		return models.ChildOrder{}, fmt.Errorf("invalid side: %s", side)
	}

	if timeInForce == "" {
		return models.ChildOrder{}, fmt.Errorf("time in force is required")
	}

	if timeInForce != TimeInForceGTC && timeInForce != TimeInForceIOC && timeInForce != TimeInForceFOK {
		return models.ChildOrder{}, fmt.Errorf("invalid time in force: %s", timeInForce)
	}

	req := models.SendChildOrderRequest{
		ProductCode:    productCode,
		ChildOrderType: childOrderType,
		Side:           side,
		Price:          price,
		Size:           size,
		MinuteToExpire: minuteToExpire,
		TimeInForce:    timeInForce,
	}

	return bu.PrivateAPI.PostSendChildOrder(req, isDry)
}

func (bu *BitflyerUseCase) PostCancelChildOrder(productCode, childOrderID string, isDry bool) error {
	if productCode == "" {
		productCode = ProductCodeBTCJPY
	}
	if !validateProductCode(productCode) {
		return fmt.Errorf("invalid product code: %s", productCode)
	}

	if childOrderID == "" {
		return fmt.Errorf("child order id is required")
	}

	req := models.CancelChildOrderRequest{
		ProductCode:  productCode,
		ChildOrderID: childOrderID,
	}

	return bu.PrivateAPI.PostCancelChildOrder(req, isDry)
}

func (bu *BitflyerUseCase) GetChildOrders() ([]models.ChildOrder, error) {
	return bu.PrivateAPI.GetChildOrders()
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
