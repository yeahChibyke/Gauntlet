package nvidia

import "github.com/yeahChibyke/Gauntlet/internal/protocol/canonical"

// ToChatCompletionRequest converts a canonical request into an NVIDIA
// Chat Completions request.
func ToChatCompletionRequest(
	req *canonical.Request,
) *ChatCompletionRequest {

	messages := make([]Message, 0, len(req.Messages))

	for _, msg := range req.Messages {
		messages = append(messages, Message{
			Role:    string(msg.Role),
			Content: msg.Content,
		})
	}

	return &ChatCompletionRequest{
		Model:       req.Model,
		Messages:    messages,
		Temperature: req.Temperature,
		MaxTokens:   req.MaxTokens,
		Stream:      req.Stream,
	}
}

// FromChatCompletionResponse converts an NVIDIA Chat Completions response
// into Gauntlet's canonical response.
func FromChatCompletionResponse(
	resp *ChatCompletionResponse,
) *canonical.Response {

	out := &canonical.Response{
		Model:  resp.Model,
		Output: make([]canonical.Message, 0, len(resp.Choices)),
	}

	for _, choice := range resp.Choices {
		out.Output = append(out.Output, canonical.Message{
			Role:    canonical.Role(choice.Message.Role),
			Content: choice.Message.Content,
		})
	}

	return out
}
