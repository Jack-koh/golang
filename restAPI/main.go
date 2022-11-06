package main

import (
	"net/http"
	"restAPI/app"
)

func main() {
	http.ListenAndServe(":3000", app.NewHandler())
}
