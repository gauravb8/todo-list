package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/gauravb8/todo-list/constants"
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

func InitTodoList() (TodoList, error) {
	fp, err := os.OpenFile(constants.NotesFilePath, os.O_RDONLY|os.O_CREATE, 0755)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(fp)
	if err != nil {
		return nil, err
	}

	td := TodoList{}

	json.Unmarshal(data, &td)

	return td, nil
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

func (td TodoList) SaveList() error {
	fp, err := os.OpenFile(constants.NotesFilePath, os.O_RDWR, 0755)
	if err != nil {
		return err
	}

	defer fp.Close()

	text, err := json.MarshalIndent(td, "", "\t")
	if err != nil {
		return err
	}

	_, err = fp.Write(text)
	if err != nil {
		return err
	}

	fmt.Println("New note added successfully")

	return nil
}
