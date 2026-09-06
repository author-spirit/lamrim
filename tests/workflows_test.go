package engine_test

import (
	"path/filepath"
	"testing"

	"github.com/author-spirit/lamrim/internal/workflows"
)

func TestDiscoverWorkflows(t *testing.T) {
	root := filepath.Join("..", "workflows")
	projects, err := workflows.Discover(root)
	if err != nil {
		t.Fatal(err)
	}
	if len(projects) == 0 {
		t.Fatal("expected at least one workflow project")
	}
	if projects[0].Name != "first" {
		t.Fatalf("name: %q", projects[0].Name)
	}
}

func TestLoadFirstWorkflow(t *testing.T) {
	root := filepath.Join("..", "workflows")
	project, err := workflows.Find(root, "first")
	if err != nil {
		t.Fatal(err)
	}

	graph, err := workflows.LoadGraph(project)
	if err != nil {
		t.Fatal(err)
	}
	if graph.Name != "first" {
		t.Fatalf("graph name: %q", graph.Name)
	}
	if len(graph.Nodes) != 2 {
		t.Fatalf("nodes: %d", len(graph.Nodes))
	}
}
