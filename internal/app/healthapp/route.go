package healthapp

import (
	"net/http"

	"github.com/BuildFrom/Golang-Stdlib/internal/infrastructure/sqldb"
)

func RegisterRoutes(mux *http.ServeMux, dbService sqldb.Service) {
	api := newApp(dbService)
	mux.HandleFunc("/readiness", api.Health)
}
