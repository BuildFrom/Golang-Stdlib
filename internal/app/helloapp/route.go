package helloapp

import (
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	api := newApp()
	mux.HandleFunc("/", api.HelloWorld)
}
