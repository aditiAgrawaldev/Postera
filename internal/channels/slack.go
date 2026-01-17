package channels

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/aditiAgrawaldev/postera/internal/models"
	"github.com/aditiAgrawaldev/postera/pkg/retry"
)

type SlackChannel struct {
	WebhookURL string
}

func NewSlackChannel(webhookURL string) *SlackChannel {
	return &SlackChannel{
		WebhookURL: webhookURL,
	}
}

type SlackMessage struct {
	Text string `json:"text"`
}

func (s *SlackChannel) Send(notification models.Notification) error {
	retryConfig := retry.DefaultConfig()
	return retry.Retry(retryConfig, func() error {
		return s.sendSlackMessage(notification)
	})
}

func (s *SlackChannel) sendSlackMessage(notification models.Notification) error {
	message := SlackMessage{
		Text: fmt.Sprintf("*%s*\n\n%s",
			notification.Subject,
			notification.Body,
		),
	}

	jsonData, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Post(s.WebhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to send to Slack: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Slack API returned status: %d", resp.StatusCode)
	}

	return nil
}