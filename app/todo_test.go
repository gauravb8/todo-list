package app

import (
	"testing"
	"time"
)

func TestNewNote(t *testing.T) {
	message := "Hello There!"

	now := time.Now()

	n := NewNote(message)

	if n.Message != message {
		t.Errorf("Error. Expected: %s, Received: %s", message, n.Message)
	}

	if n.CreatedAt.Before(now) || n.CreatedAt.After(time.Now()) {
		t.Errorf("Unexpected creation time for note")
	}
}
