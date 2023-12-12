package main

import (
	say "github.com/Eirik-Haukeland/task_master/src/model"
	http_server "github.com/Eirik-Haukeland/task_master/src/view/http-server"
)

func main() {
	say.Hello("")

	err := say.StoreValue("userName", "Eirik")
	if err != nil {
		panic(err)
	}

	err, name := say.GetValue("userName")

	if err != nil {
		panic(err)
	} else {
		say.Hello(name)
	}

	http_server.HttpServer()
}
