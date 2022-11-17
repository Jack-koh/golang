package main

import (
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"todo/app"
)

func main() {
	m := app.MakeNewHandler()
	n := negroni.Classic()
	n.UseHandler(m)

	log.Println("Started App")
	err := http.ListenAndServe(":3000", n)

	if err != nil {
		panic(err)
	}
}
