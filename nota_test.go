package main

import (
	"context"
	"testing"
)

func TestFindNote(t *testing.T) {
	t.Run("when there's no database", func(t *testing.T) {
		t.Run("it returns a NoDatabaseError", func(t *testing.T) {
			if _, err := FindNote(context.TODO(), db, "123"); err.(type) != NoDatabaseError {
				t.Fatal("expected a NoDatabaseError, got a %t", err)
			}
		})
	})
}
