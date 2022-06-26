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
	if req.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Welcome to todo list!")
}

func addNote(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

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
	if req.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(tdl.GetListJson())
}

func main() {

	http.HandleFunc("/", hello)
	http.HandleFunc("/list", showList)
	http.HandleFunc("/list/add", addNote)

	http.ListenAndServe(":8090", nil)

}
