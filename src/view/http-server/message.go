package http_server

// look for better name
type newTodo struct {
	Text string `json:"text"`
}

type returnTodo struct {
	Text string `json:"text"`
	id   string `json:"id"`
}
