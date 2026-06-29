package provider

import (
	"context"

	"github.com/yeahChibyke/Gauntlet/internal/protocol/canonical"
)

// StreamReader represents a provider-agnostic stream of canonical responses.
type StreamReader interface {
	Next(ctx context.Context) (*canonical.Response, bool, error)
	Close() error
}
