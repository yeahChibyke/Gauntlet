package provider

import (
	"context"

	"github.com/yeahChibyke/Gauntlet/internal/protocol/canonical"
	"github.com/yeahChibyke/Gauntlet/internal/protocol/responses"
)

// Provider defines the contract every backend (NVIDIA, OpenRouter,
// Ollama, OpenAI, etc.) must implement.
type Provider interface {
	Responses(
		ctx context.Context,
		req *canonical.Request,
	) (*responses.Response, error)
}
