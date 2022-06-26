package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gauravb8/todo-list/app"
)

type AddNoteReq struct {
	Message string
}

var tdl = app.TodoList{}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Welcome to todo list!")
}

func addNote(w http.ResponseWriter, req *http.Request) {
	c := req.Context()

	fmt.Printf("%s: got /hello request\n", c.Value("keyServerAddr"))

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("Body: ", string(body))

	postReq := AddNoteReq{}
	err = json.Unmarshal(body, &postReq)

	if err != nil {
		fmt.Println("Error while unmarshaling json:", err)
	}

	tdl = append(tdl, app.NewNote(postReq.Message))

}

func showList(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, tdl.PrintList())
}

func main() {

	http.HandleFunc("/", hello)
	http.HandleFunc("/list/add", addNote)
	http.HandleFunc("/list", showList)

	http.ListenAndServe(":8090", nil)

}
