package provider

import (
	"context"

	"github.com/yeahChibyke/Gauntlet/internal/protocol/canonical"
)

// Provider represents an LLM backend capable of executing a canonical
// request and returning a canonical response.
type Provider interface {
	Responses(
		ctx context.Context,
		req *canonical.Request,
	) (*canonical.Response, error)

	ResponsesStream(
		ctx context.Context,
		req *canonical.Request,
	) (StreamReader, error)
}
