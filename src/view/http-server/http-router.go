package http_server

import (
	"fmt"
	"io"
	"net/http"
)

func Router() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/auth", handleAuth)
	mux.HandleFunc("/api/todo", handleTodo)

	return mux
}

func handleTodo(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		fmt.Printf("/api/todo got a POST requset\n")
		io.WriteString(w, "POST request\n")

	case http.MethodPut:
		fmt.Printf("/api/todo got a PUT requset\n")
		io.WriteString(w, "PUT request\n")

	case http.MethodDelete:
		fmt.Printf("/api/todo got a Delete requset\n")
		io.WriteString(w, "Delete request\n")

	case http.MethodGet:
		fmt.Printf("/api/todo got a get requset\n")
		io.WriteString(w, "GET request\n")

	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func handleAuth(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		fmt.Printf("/api/auth got a POST requset\n")
		io.WriteString(w, "POST request\n")

	case http.MethodPut:
		fmt.Printf("/api/auth got a PUT requset\n")
		io.WriteString(w, "PUT request\n")

	case http.MethodDelete:
		fmt.Printf("/api/auth got a Delete requset\n")
		io.WriteString(w, "Delete request\n")

	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}
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
