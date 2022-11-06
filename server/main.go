package main

import (
	"golang/app"
	"net/http"
)

func main() {
	http.ListenAndServe(":3000", app.NewHttpHandler())
}
