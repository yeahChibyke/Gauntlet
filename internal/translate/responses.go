package translate

import (
	"github.com/yeahChibyke/Gauntlet/internal/protocol/canonical"
	"github.com/yeahChibyke/Gauntlet/internal/protocol/responses"
)

// ToResponses converts a canonical response into an OpenAI Responses API
// response.
func ToResponses(
	resp *canonical.Response,
) *responses.Response {

	out := &responses.Response{
		Object: "response",
		Model:  resp.Model,
		Output: make([]responses.ResponseOutput, 0, len(resp.Output)),
	}

	for _, message := range resp.Output {
		out.Output = append(out.Output, responses.ResponseOutput{
			Type:   "message",
			Status: "completed",
			Role:   string(message.Role),
			Content: []responses.ResponseContent{
				{
					Type: "output_text",
					Text: message.Content,
				},
			},
		})
	}

	return out
}
