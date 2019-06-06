package models

type (
	// ShoutInput holds the source and target image to compare
	ShoutInput struct {
		Source  []byte `json:"source"`
		Target  []byte `json:"target"`
		ShoutID string `json:"shout_id"`
	}
)
