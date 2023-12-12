package http_server

import (
	"encoding/json"
	"fmt"
	say "github.com/Eirik-Haukeland/task_master/src/model"
	"net/http"
)

func Router() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/todo/create", createTodo)

	return mux
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("/api/todo got a POST requset\n")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newTodoItem := &newTodo{}
	requestBuffer := []byte(r.Body) // todo: how to get the data formated so []byte() can access it

	err := json.Unmarshal(requestBuffer, newTodoItem)

	if err != nil {
		fmt.Printf("error unmarshaling JSON: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	todoText := newTodoItem.Text

	key := say.RandomKey()
	say.StoreValue(key, todoText)
}

/*
* reference form a tutorial
*

func getRoot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	hasFirst := r.URL.Query().Has("first")
	first := r.URL.Query().Get("first")
	hasSecond := r.URL.Query().Has("second")
	second := r.URL.Query().Get("second")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("could not read body: %s\n", err)
	}

	fmt.Printf("%s: got / request. first(%t)=%s, second(%t)=%s, body:\n%s\n",
		ctx.Value(keyServerAddr),
		hasFirst, first,
		hasSecond, second,
		body)
	io.WriteString(w, "this is my website!\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Printf("%s: got /hello requset\n", ctx.Value(keyServerAddr))

	myName := r.PostFormValue("myName")
	if myName == "" {
		w.Header().Set("x-missing-field", "myName")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	io.WriteString(w, fmt.Sprintf("hello, %s!\n", myName))
}

*/
