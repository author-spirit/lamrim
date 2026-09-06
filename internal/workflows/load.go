package workflows

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/author-spirit/lamrim/internal/engine"
)

type Config struct {
	Name        string              `json:"name"`
	Description string              `json:"description,omitempty"`
	Nodes       []NodeConfig        `json:"nodes"`
	Edges       map[string][]string `json:"edges,omitempty"`
}

type NodeConfig struct {
	ID    string          `json:"id"`
	Name  string          `json:"name"`
	Type  string          `json:"type"`
	Input json.RawMessage `json:"input,omitempty"`
}

func LoadConfig(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return Config{}, fmt.Errorf("config json: %w", err)
	}
	if cfg.Name == "" {
		return Config{}, fmt.Errorf("config name is required")
	}
	if len(cfg.Nodes) == 0 {
		return Config{}, fmt.Errorf("config nodes are required")
	}
	return cfg, nil
}

func LoadGraph(project Project) (*engine.Graph, error) {
	cfg, err := LoadConfig(filepath.Join(project.Dir, ConfigFile))
	if err != nil {
		return nil, err
	}

	graph := engine.NewGraph(cfg.Name)
	for _, nodeCfg := range cfg.Nodes {
		nodeType := engine.NodeType(nodeCfg.Type)
		if nodeCfg.Name == "" {
			return nil, fmt.Errorf("node name is required")
		}
		if nodeCfg.Type == "" {
			return nil, fmt.Errorf("node %q type is required", nodeCfg.Name)
		}

		id := nodeCfg.ID
		if id == "" {
			id = nodeCfg.Name
		}

		if _, err := graph.AddConfiguredNode(id, nodeCfg.Name, nodeType, nodeCfg.Input); err != nil {
			return nil, err
		}
	}

	for from, targets := range cfg.Edges {
		if _, ok := graph.Nodes[from]; !ok {
			return nil, fmt.Errorf("edge source %q not found", from)
		}
		for _, to := range targets {
			if _, ok := graph.Nodes[to]; !ok {
				return nil, fmt.Errorf("edge target %q not found", to)
			}
		}
		graph.Edges[from] = append(graph.Edges[from], targets...)
	}

	return graph, nil
}

func Run(project Project, graph *engine.Graph) error {
	if err := graph.Execute(); err != nil {
		return err
	}
	return RunScript(project, graph)
}
