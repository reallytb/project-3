package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	logger     *log.Logger
	httpServer *http.Server
}

func NewServer(logger *log.Logger) *Server {
	router := NewRouter(logger)
	return &Server{
		logger: logger,
		httpServer: &http.Server{
			Addr:         ":8080",
			Handler:      router,
			ErrorLog:     logger,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  15 * time.Second,
		},
	}
}

func NewRouter(logger *log.Logger) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handlers.MainHandler)
	mux.HandleFunc("POST /upload", handlers.UploadHandler)
	return mux
}

func (s *Server) Start() error {
	s.logger.Printf("Сервер запускается по адресу: %s", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}
