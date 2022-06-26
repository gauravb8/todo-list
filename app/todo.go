package app

import (
	"fmt"
	"time"
)

type Note struct {
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"createdAt"`
}

type TodoList []*Note

func NewNote(msg string) *Note {
	return &Note{
		Message:   msg,
		CreatedAt: time.Now(),
	}
}

// func (td TodoList) GetListJson() []map[string]string {
// 	var m []map[string]string
// 	for _, note := range td {
// 		temp := make(map[string]string)
// 		temp["message"] = note.Message
// 		temp["createdAt"] = note.CreatedAt.Format("02 Jan, 2006 15:04:05")
// 		m = append(m, temp)
// 	}

// 	return m
// }

func (td TodoList) PrintList() string {
	s := ""
	for i, note := range td {
		s += fmt.Sprintf("%d.\t%s\t%s\n", i+1, note.Message, note.CreatedAt.Format("02 Jan, 2006 15:04:05"))
	}

	return s
}
