package main

import (
	"net/http"
	"server/app"
)

func main() {
	http.ListenAndServe(":3000", app.NewHttpHandler())
}
