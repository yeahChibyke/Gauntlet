package responses

import (
	"errors"
	"fmt"

	"github.com/yeahChibyke/Gauntlet/internal/protocol/canonical"
)

var (
	ErrUnsupportedContentType = errors.New("unsupported content type")
	ErrEmptyMessage           = errors.New("message contains no text content")
)

// ToCanonical converts an OpenAI Responses request into
// Gauntlet's canonical request representation.
func (r Request) ToCanonical() (*canonical.Request, error) {
	req := &canonical.Request{
		Model:       r.Model,
		Temperature: r.Temperature,
		MaxTokens:   r.MaxOutputTokens,
		Stream:      r.Stream,
	}

	for _, msg := range r.Input {
		canonicalMsg := canonical.Message{
			Role: canonical.Role(msg.Role),
		}

		for _, part := range msg.Content {
			switch part.Type {
			case "input_text", "output_text", "text":
				if part.Text != "" {
					if canonicalMsg.Content != "" {
						canonicalMsg.Content += "\n"
					}
					canonicalMsg.Content += part.Text
				}

			default:
				return nil, fmt.Errorf("%w: %s", ErrUnsupportedContentType, part.Type)
			}
		}

		if canonicalMsg.Content == "" {
			return nil, ErrEmptyMessage
		}

		req.Messages = append(req.Messages, canonicalMsg)
	}

	return req, nil
}
