package handler

import (
	"net/http"
	"strconv"

	"github.com/bitcoin-trading-automation/internal/bitflyer-api/api/models"
	"github.com/bitcoin-trading-automation/internal/config"
	"github.com/bitcoin-trading-automation/internal/usecase"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Config  config.Config
	UseCase usecase.ITickerLog
}

type Ihandler interface {
	GetTickerLogs(ctx *gin.Context)
	GetTickerLogByTickID(ctx *gin.Context)
	PostTickerLog(ctx *gin.Context)
}

func NewHandler(cfg config.Config) (Ihandler, error) {
	useCase, err := usecase.NewTickerLog(cfg)
	if err != nil {
		return nil, err
	}

	return &Handler{
		Config:  cfg,
		UseCase: useCase,
	}, nil
}

func (h *Handler) GetTickerLogs(ctx *gin.Context) {
	tickerLogs, err := h.UseCase.GetTickerLogs()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tickerLogs)
}

func (h *Handler) GetTickerLogByTickID(ctx *gin.Context) {
	tickerIDStr := ctx.Param("tickerID")
	tickerID, err := strconv.Atoi(tickerIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tickerID"})
		return
	}

	ticker, err := h.UseCase.GetTickerLogByTickID(tickerID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if ticker == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Ticker not found"})
		return
	}

	ctx.JSON(http.StatusOK, ticker)
}

func (h *Handler) PostTickerLog(ctx *gin.Context) {
	var ticker models.Ticket
	if err := ctx.BindJSON(&ticker); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.UseCase.PostTickerLog(ticker)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Ticker log posted"})
}
