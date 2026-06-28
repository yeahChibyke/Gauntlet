package nvidia

import (
	"context"

	"github.com/yeahChibyke/Gauntlet/internal/protocol/canonical"
)

// Provider implements the generic provider.Provider interface
// using NVIDIA NIM.
type Provider struct {
	client *Client
}

// NewProvider constructs an NVIDIA provider.
func NewProvider() (*Provider, error) {
	client, err := New()
	if err != nil {
		return nil, err
	}

	return &Provider{
		client: client,
	}, nil
}

// Responses executes a canonical request against NVIDIA NIM.
func (p *Provider) Responses(
	ctx context.Context,
	req *canonical.Request,
) (*canonical.Response, error) {

	nvidiaReq := ToChatCompletionRequest(req)

	resp, err := p.client.ChatCompletion(ctx, nvidiaReq)
	if err != nil {
		return nil, err
	}

	return FromChatCompletionResponse(resp), nil
}
