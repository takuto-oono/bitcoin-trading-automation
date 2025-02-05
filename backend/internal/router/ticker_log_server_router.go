package router

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"

	"github.com/bitcoin-trading-automation/internal/config"
	"github.com/bitcoin-trading-automation/internal/handler"
)

func NewTickerLogServerRouter(cfg config.Config) (*gin.Engine, error) {
	h, err := handler.NewTickerLogHandler(cfg)
	if err != nil {
		return nil, err
	}

	router := gin.Default()

	router.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	router.GET("/ticker-logs", h.GetTickerLogs)
	router.GET("/ticker-logs/:tickerID", h.GetTickerLogByTickID)
	router.POST("/ticker-logs", h.PostTickerLog)

	return router, nil
}

func RunTickerLogServer(r *gin.Engine, cfg config.Config) error {
	u, err := url.Parse(cfg.Url.TickerLogServer)
	if err != nil {
		return err
	}

	return r.Run(":" + u.Port())
}
