package models

import "errors"

type SlackNotificationPostMessage struct {
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

func (b *SlackNotificationPostMessage) Validate() error {
	if b.Channel == "" {
		return errors.New("channel is required")
	}

	if b.Text == "" {
		return errors.New("text is required")
	}

	return nil
}
