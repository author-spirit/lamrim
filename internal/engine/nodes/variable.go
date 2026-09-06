package nodes

import (
	"encoding/json"
	"fmt"
)

// Variable is the planned set/delete work for a variable node.
type Variable struct {
	Set    map[string]any `json:"set,omitempty"`
	Delete []string       `json:"delete,omitempty"`
}

func parseVariable(input json.RawMessage) (Variable, error) {
	var spec Variable
	if err := decodeInput(input, &spec); err != nil {
		return Variable{}, fmt.Errorf("variable input: %w", err)
	}
	if spec.Set == nil {
		spec.Set = make(map[string]any)
	}
	return spec, nil
}

func ExecuteVariable(input json.RawMessage, c Context) (map[string]any, error) {
	if c == nil {
		return nil, fmt.Errorf("nil context")
	}

	spec, err := parseVariable(input)
	if err != nil {
		return nil, err
	}

	set := make(map[string]any, len(spec.Set))
	for key, value := range spec.Set {
		set[key] = value
		c.SetVariable(key, value)
	}
	for _, key := range spec.Delete {
		c.DeleteVariable(key)
	}

	return set, nil
}

func SetVariableKey(input json.RawMessage, key string, value any) (json.RawMessage, error) {
	spec, err := parseVariable(input)
	if err != nil {
		return nil, err
	}
	spec.Set[key] = value
	spec.Delete = removeKey(spec.Delete, key)
	return encodeInput(spec)
}

func SetVariableKeys(input json.RawMessage, raw string) (json.RawMessage, error) {
	var values map[string]any
	if err := json.Unmarshal([]byte(raw), &values); err != nil {
		return nil, fmt.Errorf("variables json: %w", err)
	}

	spec, err := parseVariable(input)
	if err != nil {
		return nil, err
	}
	for key, value := range values {
		spec.Set[key] = value
		spec.Delete = removeKey(spec.Delete, key)
	}
	return encodeInput(spec)
}

func DeleteVariableKey(input json.RawMessage, key string) (json.RawMessage, error) {
	spec, err := parseVariable(input)
	if err != nil {
		return nil, err
	}
	delete(spec.Set, key)
	if !containsKey(spec.Delete, key) {
		spec.Delete = append(spec.Delete, key)
	}
	return encodeInput(spec)
}

func GetVariableKey(input json.RawMessage, key string) (any, bool) {
	spec, err := parseVariable(input)
	if err != nil {
		return nil, false
	}
	value, ok := spec.Set[key]
	return value, ok
}

func removeKey(keys []string, key string) []string {
	out := keys[:0]
	for _, k := range keys {
		if k != key {
			out = append(out, k)
		}
	}
	return out
}

func containsKey(keys []string, key string) bool {
	for _, k := range keys {
		if k == key {
			return true
		}
	}
	return false
}
