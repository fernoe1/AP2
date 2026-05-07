package mailjet

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/fernoe1/AP2/assignment-1/notification/internal/domain"
)

type Client struct {
	apiKey     string
	apiSecret  string
	fromEmail  string
	fromName   string
	mode       string
	httpClient *http.Client
}

func InitClient() *Client {
	return &Client{
		apiKey:     os.Getenv("MAILJET_API_KEY"),
		apiSecret:  os.Getenv("MAILJET_API_SECRET"),
		fromEmail:  envOr("MAILJET_FROM_EMAIL", "krutoytemirlan2007@gmail.com"),
		fromName:   envOr("MAILJET_FROM_NAME", "test"),
		httpClient: &http.Client{Timeout: 5 * time.Second},
	}
}

func envOr(name, fallback string) string {
	v := os.Getenv(name)
	if v == "" {
		return fallback
	}
	return v
}

type sendPayload struct {
	Messages []struct {
		From struct {
			Email string `json:"Email"`
			Name  string `json:"Name"`
		} `json:"From"`
		To []struct {
			Email string `json:"Email"`
			Name  string `json:"Name,omitempty"`
		} `json:"To"`
		Subject  string `json:"Subject"`
		TextPart string `json:"TextPart"`
	} `json:"Messages"`
}

type sendResponse struct {
	Messages []struct {
		To []struct {
			MessageUUID string `json:"MessageUUID"`
			MessageID   int64  `json:"MessageID"`
		} `json:"To"`
	} `json:"Messages"`
}

func (c *Client) Send(ctx context.Context, notification *domain.Notification) error {
	if c.mode == "stub" || c.apiKey == "" || c.apiSecret == "" {
		return nil
	}

	fmt.Println("yo")

	payload := sendPayload{
		Messages: []struct {
			From struct {
				Email string `json:"Email"`
				Name  string `json:"Name"`
			} `json:"From"`
			To []struct {
				Email string `json:"Email"`
				Name  string `json:"Name,omitempty"`
			} `json:"To"`
			Subject  string `json:"Subject"`
			TextPart string `json:"TextPart"`
		}{
			{
				To: []struct {
					Email string `json:"Email"`
					Name  string `json:"Name,omitempty"`
				}{{Email: notification.CustomerEmail}},
				Subject:  notification.Status,
				TextPart: notification.Status,
			},
		},
	}
	payload.Messages[0].From.Email = c.fromEmail
	payload.Messages[0].From.Name = c.fromName

	raw, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://api.mailjet.com/v3.1/send", bytes.NewReader(raw))
	if err != nil {
		return err
	}

	req.SetBasicAuth(c.apiKey, c.apiSecret)
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode >= 500 {
		return fmt.Errorf("mailjet temporary error: %d", res.StatusCode)
	}
	if res.StatusCode >= 400 {
		return fmt.Errorf("mailjet permanent error: %d", res.StatusCode)
	}

	var out sendResponse
	if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
		return err
	}
	if len(out.Messages) == 0 || len(out.Messages[0].To) == 0 {
		return fmt.Errorf("mailjet response missing message id")
	}

	fmt.Println(out)

	if out.Messages[0].To[0].MessageUUID != "" {
		return nil
	}
	return nil
}
