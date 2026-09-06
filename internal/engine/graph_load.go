package engine

import (
	"encoding/json"
	"fmt"
)

// AddConfiguredNode adds a node with an optional stable ID from workflow config.
func (g *Graph) AddConfiguredNode(id, name string, nodetype NodeType, input json.RawMessage) (*Node, error) {
	if id == "" {
		id = generateNodeId(nodetype)
	}
	if _, exists := g.Nodes[id]; exists {
		return nil, fmt.Errorf("duplicate node id %q", id)
	}

	node := &Node{
		ID:    id,
		Name:  name,
		Type:  nodetype,
		Input: input,
	}

	g.Nodes[id] = node
	g.Edges[id] = []string{}
	g.nodeOrder = append(g.nodeOrder, id)
	return node, nil
}
