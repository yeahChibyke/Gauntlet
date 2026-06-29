package service

import (
	"context"

	"github.com/yeahChibyke/Gauntlet/internal/protocol/responses"
	"github.com/yeahChibyke/Gauntlet/internal/provider"
	"github.com/yeahChibyke/Gauntlet/internal/translate"
)

// ResponseService coordinates handling of OpenAI Responses requests.
type ResponseService struct {
	provider provider.Provider
}

// NewResponseService constructs a ResponseService.
func NewResponseService(
	provider provider.Provider,
) *ResponseService {
	return &ResponseService{
		provider: provider,
	}
}

// Handle converts an OpenAI Responses request into a canonical request,
// executes it against the configured provider, then converts the result
// back into an OpenAI Responses response.
func (s *ResponseService) Handle(
	ctx context.Context,
	req *responses.Request,
) (*responses.Response, error) {

	canonicalReq, err := req.ToCanonical()
	if err != nil {
		return nil, err
	}

	canonicalResp, err := s.provider.Responses(
		ctx,
		canonicalReq,
	)
	if err != nil {
		return nil, err
	}

	return translate.ToResponses(canonicalResp), nil
}

func (s *ResponseService) HandleStream(
	ctx context.Context,
	req *responses.Request,
) (provider.StreamReader, error) {

	canonicalReq, err := req.ToCanonical()
	if err != nil {
		return nil, err
	}

	return s.provider.ResponsesStream(
		ctx,
		canonicalReq,
	)
}