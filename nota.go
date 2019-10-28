package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

// A Note is just a decorated string with a title and some attributes.
type Note struct {
	ID, Title, Content, ContentType string
	CreatedAt, UpdatedAt            time.Time
}

// ErrNotFound is an error indicating that a record wasn't found
type ErrNotFound struct {
	Table, ID string
}

// Error gets the error message for an ErrNotFound
func (e ErrNotFound) Error() string {
	return fmt.Printf("Not found: table=%q, row=%q", e.Table, e.ID)
}

// FindNote Finds a Note by ID in the database db
func FindNote(ctx context.Context, db sql.DB, id string) (*Note, error) {
	q := "SELECT * FROM notes WHERE id = ?"
	note := &Note{}
	err := db.QueryRowContext(ctx, id).Scan(&note)
	return note, err
}
