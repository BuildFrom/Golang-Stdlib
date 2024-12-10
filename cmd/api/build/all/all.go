package all

import (
	"net/http"

	"github.com/BuildFrom/Golang-Stdlib/internal/app/healthapp"
	"github.com/BuildFrom/Golang-Stdlib/internal/app/helloapp"
	"github.com/BuildFrom/Golang-Stdlib/internal/app/todoapp"
	"github.com/BuildFrom/Golang-Stdlib/internal/infrastructure/sqldb"
)

// RegisterRoutes registers all routes for the API.
func RegisterRoutes(dbService sqldb.Service) http.Handler {
	mux := http.NewServeMux()

	helloapp.RegisterRoutes(mux)
	healthapp.RegisterRoutes(mux, dbService)
	todoapp.RegisterRoutes(mux, dbService)

	return mux
}
