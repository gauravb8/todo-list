package app

import (
	"testing"
	"time"
)

func TestNewNote(t *testing.T) {
	message := "Hello There!"

	now := time.Now()

	n := NewNote(message)

	if n.message != message {
		t.Errorf("Error. Expected: %s, Received: %s", message, n.message)
	}

	if n.createdAt.Before(now) || n.createdAt.After(time.Now()) {
		t.Errorf("Unexpected creation time for note")
	}
}
