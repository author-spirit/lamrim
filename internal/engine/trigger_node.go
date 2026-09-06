package engine

import (
	"fmt"

	"github.com/author-spirit/lamrim/internal/engine/nodes"
)

func (n *Node) SetTrigger(kind string) error {
	if n == nil {
		return fmt.Errorf("nil node")
	}
	if n.Type != Trigger {
		return n.wrongType(Trigger)
	}
	input, err := nodes.SetTriggerType(n.Input, kind)
	if err != nil {
		return err
	}
	n.Input = input
	return nil
}
