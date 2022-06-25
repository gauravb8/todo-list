package app

import "time"

type Note struct {
	message   string
	createdAt time.Time
}

type TodoList []*Note

func NewNote(msg string) *Note {
	return &Note{
		message:   msg,
		createdAt: time.Now(),
	}
}
