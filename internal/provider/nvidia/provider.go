package nvidia

import (
	"context"

	"github.com/yeahChibyke/Gauntlet/internal/config"
	"github.com/yeahChibyke/Gauntlet/internal/protocol/canonical"
)

type Provider struct {
	client *Client
}

func NewProvider(cfg *config.Config) (*Provider, error) {
	client, err := New(cfg)
	if err != nil {
		return nil, err
	}

	return &Provider{
		client: client,
	}, nil
}

func (p *Provider) Responses(
	ctx context.Context,
	req *canonical.Request,
) (*canonical.Response, error) {

	nvidiaReq := ToChatCompletionRequest(req)

	nvidiaResp, err := p.client.ChatCompletion(ctx, nvidiaReq)
	if err != nil {
		return nil, err
	}

	return FromChatCompletionResponse(nvidiaResp), nil
}
