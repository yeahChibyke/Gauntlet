package service

import (
	"github.com/yeahChibyke/Gauntlet/internal/protocol/canonical"
	"github.com/yeahChibyke/Gauntlet/internal/protocol/responses"
)

// ResponseService coordinates handling of OpenAI Responses API requests.
type ResponseService struct{}

// NewResponseService constructs a ResponseService.
func NewResponseService() *ResponseService {
	return &ResponseService{}
}

// Handle converts an OpenAI Responses request into Gauntlet's
// canonical request representation.
func (s *ResponseService) Handle(
	req *responses.Request,
) (*canonical.Request, error) {
	return req.ToCanonical()
}
