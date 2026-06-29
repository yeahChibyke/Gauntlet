package nvidia

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

// StreamReader reads NVIDIA Server-Sent Events.
type StreamReader struct {
	scanner *bufio.Scanner
	closer  io.Closer
}

// NewStreamReader constructs a StreamReader.
func NewStreamReader(r io.ReadCloser) *StreamReader {
	scanner := bufio.NewScanner(r)

	// Allow larger streamed tokens than Scanner's default 64 KiB.
	const maxCapacity = 1024 * 1024
	scanner.Buffer(make([]byte, 4096), maxCapacity)

	return &StreamReader{
		scanner: scanner,
		closer:  r,
	}
}

// Close releases the underlying HTTP body.
func (s *StreamReader) Close() error {
	return s.closer.Close()
}

// Next returns the next streamed chunk.
//
// ok == false means the stream has ended normally.
func (s *StreamReader) Next(
	ctx context.Context,
) (*ChatCompletionChunk, bool, error) {

	for {

		select {
		case <-ctx.Done():
			return nil, false, ctx.Err()
		default:
		}

		if !s.scanner.Scan() {

			if err := s.scanner.Err(); err != nil {
				return nil, false, err
			}

			return nil, false, nil
		}

		line := strings.TrimSpace(s.scanner.Text())

		if line == "" {
			continue
		}

		if !strings.HasPrefix(line, "data: ") {
			continue
		}

		payload := strings.TrimPrefix(line, "data: ")

		if payload == "[DONE]" {
			return nil, false, nil
		}

		var chunk ChatCompletionChunk

		if err := json.Unmarshal([]byte(payload), &chunk); err != nil {
			return nil, false, fmt.Errorf(
				"decode stream chunk: %w",
				err,
			)
		}

		return &chunk, true, nil
	}
}
