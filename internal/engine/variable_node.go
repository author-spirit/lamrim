package engine

import (
	"fmt"

	"github.com/author-spirit/lamrim/internal/engine/nodes"
)

func (n *Node) SetVariable(key string, value any) error {
	if n == nil {
		return fmt.Errorf("nil node")
	}
	if n.Type != Variable {
		return n.wrongType(Variable)
	}
	input, err := nodes.SetVariableKey(n.Input, key, value)
	if err != nil {
		return err
	}
	n.Input = input
	return nil
}

func (n *Node) SetVariables(raw string) error {
	if n == nil {
		return fmt.Errorf("nil node")
	}
	if n.Type != Variable {
		return n.wrongType(Variable)
	}
	input, err := nodes.SetVariableKeys(n.Input, raw)
	if err != nil {
		return err
	}
	n.Input = input
	return nil
}

func (n *Node) DeleteVariable(key string) error {
	if n == nil {
		return fmt.Errorf("nil node")
	}
	if n.Type != Variable {
		return n.wrongType(Variable)
	}
	input, err := nodes.DeleteVariableKey(n.Input, key)
	if err != nil {
		return err
	}
	n.Input = input
	return nil
}

// GetVariable returns a value this node is configured to set, not graph runtime state.
func (n *Node) GetVariable(key string) (any, bool) {
	if n == nil || n.Type != Variable {
		return nil, false
	}
	return nodes.GetVariableKey(n.Input, key)
}
