package engine

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
)

type Graph struct {
	Name      string              `json:"name"`
	Nodes     map[string]*Node    `json:"nodes"`
	Edges     map[string][]string `json:"edges"`
	Context   *Context            `json:"context"`
	nodeOrder []string
}

func NewGraph(name string) *Graph {
	return &Graph{
		Name:    name,
		Nodes:   make(map[string]*Node),
		Edges:   make(map[string][]string),
		Context: NewContext(),
	}
}

func generateNodeId(nodeType NodeType) string {
	const text = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const ln = 8
	b := make([]byte, ln)

	for i := range b {
		b[i] = text[rand.IntN(len(text))]
	}

	return string(nodeType) + "_" + string(b)
}

func (g *Graph) AddNode(name string, nodetype NodeType) *Node {
	nodeId := generateNodeId(nodetype)
	node := &Node{
		ID:   nodeId,
		Type: nodetype,
		Name: name,
	}

	g.Nodes[nodeId] = node
	g.Edges[nodeId] = []string{}
	g.nodeOrder = append(g.nodeOrder, nodeId)
	return node
}

func (g *Graph) RepresentGraph() string {
	encoded, err := json.Marshal(g)
	if err != nil {
		return ""
	}
	return string(encoded)
}

// Execute runs each node and stores its result on the graph context.
func (g *Graph) Execute() error {
	if g.Context == nil {
		g.Context = NewContext()
	}

	for _, id := range g.nodeOrder {
		node := g.Nodes[id]
		result, err := node.ExecuteNode(g.Context)
		if err != nil {
			return fmt.Errorf("node %s: %w", id, err)
		}
		g.Context.RecordStep(id, result)
	}
	return nil
}
