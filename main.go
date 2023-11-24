package main

import (
	say "github.com/Eirik-Haukeland/task_master/src/model"
)

func main() {
	say.Hello("")

	say.StoreValue("userName", "Eirik")

	err, name := say.GetValue("userName")

	if err != nil {
		panic(err)
	} else {
		say.Hello(name)
	}

}
