package nvidia

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yeahChibyke/Gauntlet/internal/config"
	"github.com/yeahChibyke/Gauntlet/internal/provider"
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

	if resp.StatusCode >= 300 {

		var providerErr ErrorResponse

		if err := json.NewDecoder(resp.Body).Decode(&providerErr); err == nil &&
			providerErr.Error.Message != "" {

			return nil, &provider.Error{
				Provider:  "nvidia",
				Status:    resp.StatusCode,
				Message:   providerErr.Error.Message,
				Retryable: resp.StatusCode >= 500,
			}
		}

		return nil, &provider.Error{
			Provider:  "nvidia",
			Status:    resp.StatusCode,
			Message:   resp.Status,
			Retryable: resp.StatusCode >= 500,
		}
	}

	return &out, nil
}

func (c *Client) ChatCompletionStream(
	ctx context.Context,
	req *ChatCompletionRequest,
) (*StreamReader, error) {

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

	httpReq.Header.Set(
		"Authorization",
		"Bearer "+c.apiKey,
	)

	httpReq.Header.Set(
		"Content-Type",
		"application/json",
	)

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 300 {
		defer resp.Body.Close()

		return nil, &provider.Error{
			Provider:  "nvidia",
			Status:    resp.StatusCode,
			Message:   resp.Status,
			Retryable: resp.StatusCode >= 500,
		}
	}

	return NewStreamReader(resp.Body), nil
}
