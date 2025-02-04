package handler

import (
	"net/http"
	"strconv"

	"github.com/bitcoin-trading-automation/internal/bitflyer-api/api/models"
	"github.com/bitcoin-trading-automation/internal/config"
	"github.com/bitcoin-trading-automation/internal/usecase"
	"github.com/gin-gonic/gin"
)

type TickerLogHandler struct {
	Config  config.Config
	UseCase usecase.ITickerLog
}

type ITickerLogHandler interface {
	GetTickerLogs(ctx *gin.Context)
	GetTickerLogByTickID(ctx *gin.Context)
	PostTickerLog(ctx *gin.Context)
}

func NewTickerLogHandler(cfg config.Config) (ITickerLogHandler, error) {
	useCase, err := usecase.NewTickerLog(cfg)
	if err != nil {
		return nil, err
	}

	return &TickerLogHandler{
		Config:  cfg,
		UseCase: useCase,
	}, nil
}

func (h *TickerLogHandler) GetTickerLogs(ctx *gin.Context) {
	tickerLogs, statusCode, err := h.UseCase.GetTickerLogs()
	if err != nil {
		ctx.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(statusCode, tickerLogs)
}

func (h *TickerLogHandler) GetTickerLogByTickID(ctx *gin.Context) {
	tickerIDStr := ctx.Param("tickerID")
	tickerID, err := strconv.Atoi(tickerIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tickerID"})
		return
	}

	ticker, statusCode, err := h.UseCase.GetTickerLogByTickID(tickerID)
	if err != nil {
		ctx.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(statusCode, ticker)
}

func (h *TickerLogHandler) PostTickerLog(ctx *gin.Context) {
	var ticker models.Ticket
	if err := ctx.BindJSON(&ticker); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	statusCode, err := h.UseCase.PostTickerLog(ticker)
	if err != nil {
		ctx.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(statusCode, gin.H{"message": "Ticker log posted"})
}
