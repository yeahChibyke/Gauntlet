package nvidia

import (
	"context"

	"github.com/yeahChibyke/Gauntlet/internal/protocol/canonical"
)

type CanonicalStreamReader struct {
	reader *StreamReader
}

func NewCanonicalStreamReader(
	reader *StreamReader,
) *CanonicalStreamReader {
	return &CanonicalStreamReader{
		reader: reader,
	}
}

func (c *CanonicalStreamReader) Close() error {
	return c.reader.Close()
}

func (c *CanonicalStreamReader) Next(
	ctx context.Context,
) (*canonical.Response, bool, error) {

	chunk, ok, err := c.reader.Next(ctx)
	if err != nil || !ok {
		return nil, ok, err
	}

	var text string

	if len(chunk.Choices) > 0 {
		text = chunk.Choices[0].Delta.Content
	}

	return &canonical.Response{
		Model: chunk.Model,
		Output: []canonical.Message{
			{
				Role:    "assistant",
				Content: text,
			},
		},
	}, true, nil
}
