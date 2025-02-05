package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/bitcoin-trading-automation/internal/config"
	"github.com/bitcoin-trading-automation/internal/handler"
)

func NewRedisServerRouter(cfg config.Config) (*gin.Engine, error) {
	h, err := handler.NewRedisServerHandler(cfg)
	if err != nil {
		return nil, err
	}

	r := gin.Default()

	r.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "healthcheck ok!",
		})
	})

	r.GET("index/:index/key/:key", h.GetRedisHandler)
	r.POST("index/:index/key/:key", h.PostRedisHandler)
	r.DELETE("index/:index/key/:key", h.DeleteRedisHandler)

	return r, nil
}
