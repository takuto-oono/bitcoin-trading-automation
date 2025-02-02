package router

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"

	"github.com/bitcoin-trading-automation/internal/config"
	"github.com/bitcoin-trading-automation/internal/handler"
)

func NewRouter(cfg config.Config) (*gin.Engine, error) {
	h, err := handler.NewSlackNotificationHandler(cfg)
	if err != nil {
		return nil, err
	}

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	router.POST("/message", h.PostMessage)

	return router, nil
}

func Run(r *gin.Engine, cfg config.Config) error {
	u, err := url.Parse(cfg.Url.SlackNotification)
	if err != nil {
		return err
	}

	return r.Run(":" + u.Port())
}
