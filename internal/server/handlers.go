package server

import (
	"encoding/json"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"

	"github.com/author-spirit/lamrim/internal/workflows"
)

type workflowSummary struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

func (s *Server) handleListWorkflows(w http.ResponseWriter, _ *http.Request) {
	projects, err := workflows.Discover(s.workflowsDir)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	summaries := make([]workflowSummary, 0, len(projects))
	for _, project := range projects {
		cfg, err := workflows.LoadConfig(filepath.Join(project.Dir, workflows.ConfigFile))
		if err != nil {
			writeError(w, http.StatusInternalServerError, err.Error())
			return
		}
		summaries = append(summaries, workflowSummary{
			Name:        project.Name,
			Description: cfg.Description,
		})
	}

	writeJSON(w, http.StatusOK, summaries)
}

func (s *Server) handleRunWorkflow(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	project, err := workflows.Find(s.workflowsDir, name)
	if err != nil {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}

	graph, err := workflows.LoadGraph(project)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := workflows.Run(project, graph); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var payload any
	if err := json.Unmarshal([]byte(graph.RepresentGraph()), &payload); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, payload)
}
