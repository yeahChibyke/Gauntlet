package models

type Response struct {
	Models []Model `json:"models"`
}

type Model struct {
	ID string `json:"id"`

	Slug string `json:"slug"`

	Object string `json:"object"`

	Created int64 `json:"created"`

	OwnedBy string `json:"owned_by"`

	DisplayName string `json:"display_name"`

	Description string `json:"description,omitempty"`
}
