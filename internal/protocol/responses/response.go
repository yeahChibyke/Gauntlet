package responses

// Response is the OpenAI Responses API object returned by Gauntlet.
//
// This is intentionally a minimal implementation for Sprint 2.
// We'll extend it later with usage, tool calls, reasoning,
// annotations, and other OpenAI fields.
type Response struct {
	ID     string           `json:"id,omitempty"`
	Object string           `json:"object"`
	Model  string           `json:"model"`
	Output []ResponseOutput `json:"output"`
}

type ResponseOutput struct {
	ID      string            `json:"id,omitempty"`
	Type    string            `json:"type"`
	Status  string            `json:"status,omitempty"`
	Role    string            `json:"role,omitempty"`
	Content []ResponseContent `json:"content"`
}

type ResponseContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
}
