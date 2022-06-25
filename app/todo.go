package app

import (
	"fmt"
	"time"
)

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
func (td TodoList) PrintList() {
	for i, note := range td {
		fmt.Println(i, note.message, note.createdAt.Format("02 Jan, 2006 15:04:05"))
	}
}
