package canonical

// Request is the provider-independent representation of a model request.
//
// Every incoming protocol is translated into this structure before
// being converted into a provider-specific request.
type Request struct {
	// Target model.
	Model string

	// Conversation history.
	Messages []Message

	// Sampling temperature.
	Temperature *float64

	// Maximum number of tokens to generate.
	MaxTokens *int

	// Whether streaming responses are requested.
	Stream bool
}
