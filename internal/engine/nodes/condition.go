package nodes

import (
	"encoding/json"
	"fmt"
)

type Condition struct {
	Name       string `json:"name,omitempty"`
	Expression string `json:"expression,omitempty"`
}

func parseCondition(input json.RawMessage) (Condition, error) {
	var spec Condition
	if err := decodeInput(input, &spec); err != nil {
		return Condition{}, fmt.Errorf("condition input: %w", err)
	}
	return spec, nil
}

func ExecuteCondition(input json.RawMessage, c Context) (map[string]any, error) {
	if c == nil {
		return nil, fmt.Errorf("nil context")
	}

	spec, err := parseCondition(input)
	if err != nil {
		return nil, err
	}

	return map[string]any{
		"expression": spec.Expression,
	}, nil
}

func SetConditionExpression(input json.RawMessage, expression string) (json.RawMessage, error) {
	spec, err := parseCondition(input)
	if err != nil {
		return nil, err
	}
	spec.Expression = expression
	return encodeInput(spec)
}
