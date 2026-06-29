package nvidia

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yeahChibyke/Gauntlet/internal/config"
)

type Client struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

func New(cfg *config.Config) (*Client, error) {
	if cfg == nil {
		return nil, fmt.Errorf("config is nil")
	}

	return &Client{
		baseURL: cfg.NVIDIA.BaseURL,
		apiKey:  cfg.NVIDIA.APIKey,
		httpClient: &http.Client{
			Timeout: cfg.HTTP.Timeout,
		},
	}, nil
}

func (c *Client) ChatCompletion(
	ctx context.Context,
	req *ChatCompletionRequest,
) (*ChatCompletionResponse, error) {

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.baseURL+"/chat/completions",
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Authorization", "Bearer "+c.apiKey)
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("nvidia returned status %s", resp.Status)
	}

	var out ChatCompletionResponse

	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}

	return &out, nil
}
