package helloapp

import (
	"encoding/json"
	"log"
	"net/http"
)

type app struct{}

func newApp() *app {
	return &app{
		//
	}
}

func (h *app) HelloWorld(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{"message": "Hello World"}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, err = w.Write(jsonResp)
	if err != nil {
		log.Printf("error writing response. Err: %v", err)
	}
}
