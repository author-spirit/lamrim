package engine

import (
	"fmt"

	"github.com/author-spirit/lamrim/internal/engine/nodes"
)

func (n *Node) SetCondition(expression string) error {
	if n == nil {
		return fmt.Errorf("nil node")
	}
	if n.Type != Condition {
		return n.wrongType(Condition)
	}
	input, err := nodes.SetConditionExpression(n.Input, expression)
	if err != nil {
		return err
	}
	n.Input = input
	return nil
}
