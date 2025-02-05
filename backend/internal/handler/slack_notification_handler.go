package handler

import (
	"log"
	"net/http"

	"github.com/bitcoin-trading-automation/internal/config"
	"github.com/bitcoin-trading-automation/internal/models"
	slackapi "github.com/bitcoin-trading-automation/internal/slack-notification/slack-api"
	"github.com/gin-gonic/gin"
)

type SlackNotificationHandler struct {
	Config   config.Config
	SlackAPI slackapi.SlackAPI
}

func NewSlackNotificationHandler(cfg config.Config) (*SlackNotificationHandler, error) {
	slackAPI, err := slackapi.NewSlackAPI(&cfg)
	if err != nil {
		return nil, err
	}

	return &SlackNotificationHandler{Config: cfg, SlackAPI: *slackAPI}, nil
}

func (h *SlackNotificationHandler) PostMessage(ctx *gin.Context) {
	var req models.SlackNotificationPostMessage
	if err := ctx.BindJSON(&req); err != nil {
		log.Printf("failed to bind json: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request bind error"})
		return
	}

	if err := req.Validate(); err != nil {
		log.Printf("failed to validate request: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if err := h.SlackAPI.PostMessage(req.Channel, req.Text); err != nil {
		log.Printf("failed to post message: %v", err)
		return
	}

	log.Printf("message posted successfully")
	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}
