package canonical

// Response is Gauntlet's provider-independent response model.
//
// Every provider (NVIDIA, OpenAI, Anthropic, Gemini, Ollama, etc.)
// returns this type. API-specific translators are responsible for
// converting between provider formats and this canonical model.
type Response struct {
	Model      string
	Output     []Message
	Usage      *Usage
	StopReason string
}

// Usage contains token accounting returned by the provider.
type Usage struct {
	InputTokens  int
	OutputTokens int
	TotalTokens  int
}
