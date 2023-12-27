package mail

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type hookData struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}

type Hook struct {
	Address      string            `yaml:"address"`
	Name         string            `yaml:"name"`
	URL          string            `yaml:"url"`
	HTMLMarkdown bool              `yaml:"html_markdown"`
	Headers      map[string]string `yaml:"headers"`
}

func (h Hook) Send(from, text string) error {
	data := hookData{
		Username: h.Name,
		Content:  fmt.Sprintf("**%s**:\n%s", from, text),
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest("POST", h.URL, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	for k, v := range h.Headers {
		req.Header.Set(k, v)
	}
	req.Header.Set("User-Agent", "mailhook/0.0.1")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK, http.StatusNoContent:
		break
	default:
		return fmt.Errorf("webhook request failed with status code %d", resp.StatusCode)
	}

	return nil
}
