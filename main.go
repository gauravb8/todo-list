package main

import (
	"fmt"

	"github.com/gauravb8/todo-list/app"
)

func main() {
	tdl := app.TodoList{}

	note := app.NewNote("Hello!")

	tdl = append(tdl, note)

	fmt.Printf("%+v", tdl)
}
