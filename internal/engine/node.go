package engine

import (
	"encoding/json"
	"fmt"

	"github.com/author-spirit/lamrim/internal/engine/nodes"
)

type NodeType string

const (
	Trigger   NodeType = "trigger"
	Variable  NodeType = "variable"
	Condition NodeType = "condition"
	Switch    NodeType = "switch"
	Loop      NodeType = "loop"
	Function  NodeType = "function"
)

type Node struct {
	ID     string          `json:"id"`
	Name   string          `json:"name"`
	Type   NodeType        `json:"type"`
	Input  json.RawMessage `json:"input,omitempty"`
	Output json.RawMessage `json:"output,omitempty"`
}

type Result struct {
	Output map[string]any `json:"output"`
}

func (n *Node) ExecuteNode(c *Context) (Result, error) {
	if n == nil {
		return Result{}, fmt.Errorf("nil node")
	}
	if c == nil {
		return Result{}, fmt.Errorf("nil context")
	}

	var (
		out map[string]any
		err error
	)
	switch n.Type {
	case Trigger:
		out, err = nodes.ExecuteTrigger(n.Input, c)
	case Variable:
		out, err = nodes.ExecuteVariable(n.Input, c)
	case Condition:
		out, err = nodes.ExecuteCondition(n.Input, c)
	default:
		return Result{}, fmt.Errorf("unknown node type %q", n.Type)
	}
	if err != nil {
		return Result{}, err
	}
	return Result{Output: out}, nil
}

func (n *Node) wrongType(want NodeType) error {
	return fmt.Errorf("node %q is %s, not %s", n.Name, n.Type, want)
}
