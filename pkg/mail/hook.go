package mail

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type hookData struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}

type Hook struct {
	Address      string `yaml:"address"`
	Name         string `yaml:"name"`
	URL          string `yaml:"url"`
	HTMLMarkdown bool   `yaml:"html_markdown"`
}

func NewHook(emailAddress, name, hookURL string) *Hook {
	return &Hook{
		Address: emailAddress,
		Name:    name,
		URL:     hookURL,
	}
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

	// Send the webhook request
	resp, err := http.Post(h.URL, "application/json", bytes.NewBuffer(payload))
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
