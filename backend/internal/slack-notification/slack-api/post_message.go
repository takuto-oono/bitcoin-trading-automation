package slackapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const PostMessageUrl = "https://slack.com/api/chat.postMessage"

// TODO: slack api側からエラーが返ってくる方法が分からないので調査が必要そう(slackにメッセージ送信が失敗した際に検知できない)
func (s *SlackAPI) PostMessage(channel, text string) error {
	payload := map[string]string{
		"channel": channel,
		"text":    text,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", PostMessageUrl, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.Config.Slack.AccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to post message: %s", resp.Status)
	}

	return nil
}
