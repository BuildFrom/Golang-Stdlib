package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/BuildFrom/Golang-Stdlib/cmd/api/build/all"
	"github.com/BuildFrom/Golang-Stdlib/internal/infrastructure/sqldb"
	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port int
	db   sqldb.Service
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port: port,
		db:   sqldb.New(),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      all.RegisterRoutes(NewServer.db),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
