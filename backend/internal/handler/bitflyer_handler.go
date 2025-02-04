package handler

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"

	"github.com/bitcoin-trading-automation/internal/config"
	"github.com/bitcoin-trading-automation/internal/usecase"
)

type BitFlyerHandler struct {
	UseCase usecase.IBitflyerUseCase
	Config  config.Config
}

type IBitFlyerHandler interface {
	// Public API
	GetBoard(ctx *gin.Context)
	GetTicker(ctx *gin.Context)
	GetExecutions(ctx *gin.Context)
	GetBoardState(ctx *gin.Context)
	GetHealth(ctx *gin.Context)

	// Private API
	GetBalance(ctx *gin.Context)
	GetCollateral(ctx *gin.Context)
	PostSendChildOrder(ctx *gin.Context)
	PostCancelChildOrder(ctx *gin.Context)
	GetChildOrders(ctx *gin.Context)
}

func NewBitFlyerHandler(cfg config.Config) IBitFlyerHandler {
	u := usecase.NewBitflyerUseCase(cfg)
	return &BitFlyerHandler{UseCase: u, Config: cfg}
}

func (h *BitFlyerHandler) GetBoard(ctx *gin.Context) {
	productCode := ctx.Request.URL.Query().Get("product_code")
	board, statusCode, err := h.UseCase.GetBoard(productCode)
	if err != nil {
		ctx.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(statusCode, board)
}

func (h *BitFlyerHandler) GetTicker(ctx *gin.Context) {
	productCode := ctx.Request.URL.Query().Get("product_code")
	ticker, statusCode, err := h.UseCase.GetTicker(productCode)
	if err != nil {
		ctx.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(statusCode, ticker)
}

type GetExecutionsQueryParams struct {
	ProductCode string `form:"product_code"`
	Count       string `form:"count"`
	Before      string `form:"before"`
	After       string `form:"after"`
}

func NewGetExecutionsQueryParams(qp url.Values) *GetExecutionsQueryParams {
	return &GetExecutionsQueryParams{
		ProductCode: qp.Get("product_code"),
		Count:       qp.Get("count"),
		Before:      qp.Get("before"),
		After:       qp.Get("after"),
	}
}

func (h *BitFlyerHandler) GetExecutions(ctx *gin.Context) {
	qp := NewGetExecutionsQueryParams(ctx.Request.URL.Query())

	executions, statusCode, err := h.UseCase.GetExecutions(qp.ProductCode, qp.Count, qp.Before, qp.After)
	if err != nil {
		ctx.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(statusCode, executions)
}

func (h *BitFlyerHandler) GetBoardState(ctx *gin.Context) {
	productCode := ctx.Request.URL.Query().Get("product_code")
	boardState, statusCode, err := h.UseCase.GetBoardState(productCode)
	if err != nil {
		ctx.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(statusCode, boardState)
}

func (h *BitFlyerHandler) GetHealth(ctx *gin.Context) {
	productCode := ctx.Request.URL.Query().Get("product_code")
	health, statusCode, err := h.UseCase.GetHealth(productCode)
	if err != nil {
		ctx.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(statusCode, health)
}

func (h *BitFlyerHandler) GetBalance(ctx *gin.Context) {
	balance, statusCode, err := h.UseCase.GetBalance()
	if err != nil {
		ctx.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(statusCode, balance)
}

func (h *BitFlyerHandler) GetCollateral(ctx *gin.Context) {
	collateral, statusCode, err := h.UseCase.GetCollateral()
	if err != nil {
		ctx.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(statusCode, collateral)
}

func (h *BitFlyerHandler) PostSendChildOrder(ctx *gin.Context) {
	req, err := NewPostSendChildOrderRequestBody(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isDry := ctx.Request.URL.Query().Get("dry_run") == "1"
	childOrder, statusCode, err := h.UseCase.PostSendChildOrder(req.ProductCode, req.ChildOrderType, req.Side, req.Price, req.Size, req.MinuteToExpire, req.TimeInForce, isDry)
	if err != nil {
		ctx.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(statusCode, childOrder)
}

type PostSendChildOrderRequestBody struct {
	ProductCode    string  `json:"product_code"`
	ChildOrderType string  `json:"child_order_type"`
	Side           string  `json:"side"`
	Price          int     `json:"price"`
	Size           float64 `json:"size"`
	MinuteToExpire int     `json:"minute_to_expire"`
	TimeInForce    string  `json:"time_in_force"`
}

func NewPostSendChildOrderRequestBody(ctx *gin.Context) (PostSendChildOrderRequestBody, error) {
	var req PostSendChildOrderRequestBody
	err := ctx.BindJSON(&req)
	return req, err
}

func (h *BitFlyerHandler) PostCancelChildOrder(ctx *gin.Context) {
	req, err := NewCancelChildOrderRequestBody(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isDry := ctx.Request.URL.Query().Get("dry_run") == "1"
	statusCode, err := h.UseCase.PostCancelChildOrder(req.ProductCode, req.ChildOrderID, isDry)
	if err != nil {
		ctx.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(statusCode, gin.H{"message": "success"})
}

type CancelChildOrderRequestBody struct {
	ProductCode  string `json:"product_code"`
	ChildOrderID string `json:"child_order_id"`
}

func NewCancelChildOrderRequestBody(ctx *gin.Context) (CancelChildOrderRequestBody, error) {
	var req CancelChildOrderRequestBody
	err := ctx.BindJSON(&req)
	return req, err
}

func (h *BitFlyerHandler) GetChildOrders(ctx *gin.Context) {
	childOrders, statusCode, err := h.UseCase.GetChildOrders()
	if err != nil {
		ctx.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(statusCode, childOrders)
}
