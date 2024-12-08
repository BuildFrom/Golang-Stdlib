package healthapp

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/shohinsan/nestly/internal/infrastructure/sqldb"
)

type app struct {
	db sqldb.Service
}

func newApp(db sqldb.Service) *app {
	return &app{db: db}
}

func (a *app) Health(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(a.db.Health())
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, err = w.Write(jsonResp)
	if err != nil {
		log.Printf("error writing response. Err: %v", err)
	}
}
