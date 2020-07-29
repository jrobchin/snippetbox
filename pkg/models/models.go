package models

import (
	"errors"
	"time"
)

// ErrNoRecord is raised when no records are found
var ErrNoRecord = errors.New("models: no matching record found")

// Snippet is the model used to store snippets
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
