package main

import (
	"log"
	"net/http"
	"todo/app"
)

func main() {
	m := app.MakeNewHandler()
	defer m.Close()

	log.Println("Started App")
	err := http.ListenAndServe(":3000", m)

	if err != nil {
		panic(err)
	}
}
