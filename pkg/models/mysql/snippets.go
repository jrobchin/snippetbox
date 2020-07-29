package mysql

import (
	"database/sql"

	"github.com/jrobchin/snippetbox/pkg/models"
)

// SnippetModel wraps a sql.DB connection
type SnippetModel struct {
	DB *sql.DB
}

// Insert inserts a snippet record
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

// Get get's the snippet from an id
func (m *SnippetModel) Get(id int) ([]*models.Snippet, error) {
	return nil, nil
}

// Latest returns the 10 most recently created snippets
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
