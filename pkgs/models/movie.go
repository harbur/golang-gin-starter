package models

// Movie movie
type Movie struct {
	// id
	ID int64 `json:"id,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`
}
