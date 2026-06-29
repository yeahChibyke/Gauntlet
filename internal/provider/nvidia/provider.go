package nvidia

import (
	"context"

	"github.com/yeahChibyke/Gauntlet/internal/config"
	"github.com/yeahChibyke/Gauntlet/internal/protocol/canonical"
	"github.com/yeahChibyke/Gauntlet/internal/provider"
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

func (p *Provider) ResponsesStream(
	ctx context.Context,
	req *canonical.Request,
) (provider.StreamReader, error) {

	nvidiaReq := ToChatCompletionRequest(req)
	nvidiaReq.Stream = true

	stream, err := p.client.ChatCompletionStream(
		ctx,
		nvidiaReq,
	)
	if err != nil {
		return nil, err
	}

	return NewCanonicalStreamReader(stream), nil
}
