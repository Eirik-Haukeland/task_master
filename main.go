package main

import (
	say "github.com/Eirik-Haukeland/task_master/src/model"
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

}
