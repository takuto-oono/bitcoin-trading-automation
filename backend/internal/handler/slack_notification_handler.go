package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/bitcoin-trading-automation/internal/config"
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

type PostMessageRequestBody struct {
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

func (h *SlackNotificationHandler) PostMessage(ctx *gin.Context) {
	var req PostMessageRequestBody
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

func (b *PostMessageRequestBody) Validate() error {
	if b.Channel == "" {
		return errors.New("channel is required")
	}

	if b.Text == "" {
		return errors.New("text is required")
	}

	return nil
}
