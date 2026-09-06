package server

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	workflowsDir string
	clientDir    string
}

func New(workflowsDir, clientDir string) *Server {
	return &Server{
		workflowsDir: workflowsDir,
		clientDir:    clientDir,
	}
}

func (s *Server) Handler() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", s.handleHealth)
	r.Route("/api", func(r chi.Router) {
		r.Get("/workflows", s.handleListWorkflows)
		r.Post("/workflows/{name}/run", s.handleRunWorkflow)
	})

	fileServer := http.StripPrefix("/", http.FileServer(http.Dir(s.clientDir)))
	r.Handle("/*", fileServer)

	return r
}

func (s *Server) handleHealth(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}
