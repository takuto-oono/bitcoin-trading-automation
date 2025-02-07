package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/bitcoin-trading-automation/internal/config"
	"github.com/bitcoin-trading-automation/internal/handler"
)

func NewBitFlyerRouter(cfg config.Config) *gin.Engine {
	h := handler.NewBitFlyerHandler(cfg)

	r := gin.Default()

	return setUpHandler(r, h)
}

func setUpHandler(r *gin.Engine, h handler.IBitFlyerHandler) *gin.Engine {
	r.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "healthcheck ok!",
		})
	})

	// public API
	r.GET("/board/", h.GetBoard)
	r.GET("/ticker/", h.GetTicker)
	r.GET("/executions/", h.GetExecutions)
	r.GET("/board_state/", h.GetBoardState)
	r.GET("/health/", h.GetHealth)

	// private API
	r.GET("/getbalance/", h.GetBalance)
	r.GET("/getcollateral/", h.GetCollateral)
	r.POST("/sendchildorder/", h.PostSendChildOrder)
	r.POST("/cancelchildorder/", h.PostCancelChildOrder)
	r.GET("/getchildorders/", h.GetChildOrders)

	return r
}
