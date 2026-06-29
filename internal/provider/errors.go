package provider

import "fmt"

// Error represents an error returned by an LLM provider.
type Error struct {
	Provider  string
	Status    int
	Message   string
	Retryable bool
}

func (e *Error) Error() string {
	return fmt.Sprintf(
		"%s provider error (%d): %s",
		e.Provider,
		e.Status,
		e.Message,
	)
}
