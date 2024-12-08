package todoapp

import (
	"net/http"

	"github.com/shohinsan/nestly/internal/infrastructure/sqldb"
	mw "github.com/shohinsan/nestly/internal/sdk/middleware"
)

func RegisterRoutes(mux *http.ServeMux, dbService sqldb.Service) http.Handler {
	repo := newStore(dbService)
	api := newApp(repo)

	mux.HandleFunc("POST /todo", api.createTodoHandler)
	mux.HandleFunc("GET /{$}", api.getTodosHandler)
	mux.HandleFunc("GET /todo/{id}", api.getTodoByIDHandler)
	mux.HandleFunc("PUT /todo/{id}", api.updateTodoHandler)
	mux.HandleFunc("DELETE /todo/{id}", api.deleteTodoHandler)

	chains := []mw.Middleware{
		mw.CORS,
	}

	return mw.WrapMiddleware(mux, chains...)
}

// dbService
// repo := newStore()

//

// mux.HandleFunc("POST /todo", api.createTodoHandler)
// mux.HandleFunc("GET /{$}", api.getTodosHandler)
// mux.HandleFunc("PUT /todo/{id}", api.updateTodoHandler)
// mux.HandleFunc("GET /todo/{id}", api.getTodoByIDHandler)
// mux.HandleFunc("DELETE /todo/{id}", api.deleteTodoHandler)
