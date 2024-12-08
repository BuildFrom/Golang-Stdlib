package all

import (
	"net/http"

	"github.com/shohinsan/nestly/internal/app/healthapp"
	"github.com/shohinsan/nestly/internal/app/helloapp"
	"github.com/shohinsan/nestly/internal/app/todoapp"
	"github.com/shohinsan/nestly/internal/infrastructure/sqldb"
)

// RegisterRoutes registers all routes for the API.
func RegisterRoutes(dbService sqldb.Service) http.Handler {
	mux := http.NewServeMux()

	helloapp.RegisterRoutes(mux)
	healthapp.RegisterRoutes(mux, dbService)
	todoapp.RegisterRoutes(mux, dbService)

	return mux
}
