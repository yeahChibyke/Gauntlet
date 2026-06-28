package responses

// Request represents the subset of the OpenAI Responses API
// that Gauntlet currently supports.
//
// Sprint 1 intentionally supports:
//
//   - model
//   - input (text messages)
//   - temperature
//   - max_output_tokens
//   - stream
//
// Additional fields (tools, reasoning, metadata, etc.) will
// be introduced incrementally as Gauntlet gains support.
type Request struct {
	Model           string    `json:"model"`
	Input           []Message `json:"input"`
	Temperature     *float64  `json:"temperature,omitempty"`
	MaxOutputTokens *int      `json:"max_output_tokens,omitempty"`
	Stream          bool      `json:"stream"`
}

type Message struct {
	Role    string    `json:"role"`
	Content []Content `json:"content"`
}

type Content struct {
	Type string `json:"type"`
	Text string `json:"text,omitempty"`
}
