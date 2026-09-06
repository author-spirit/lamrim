package nodes

import (
	"encoding/json"
	"fmt"
)

type TriggerType string

const (
	TriggerUnknown TriggerType = ""
	TimeTrigger    TriggerType = "time"
	EventTrigger   TriggerType = "event"
)

type Trigger struct {
	Type TriggerType `json:"type,omitempty"`
}

func parseTrigger(input json.RawMessage) (Trigger, error) {
	var spec Trigger
	if err := decodeInput(input, &spec); err != nil {
		return Trigger{}, fmt.Errorf("trigger input: %w", err)
	}
	return spec, nil
}

func ExecuteTrigger(input json.RawMessage, c Context) (map[string]any, error) {
	if c == nil {
		return nil, fmt.Errorf("nil context")
	}

	spec, err := parseTrigger(input)
	if err != nil {
		return nil, err
	}
	if spec.Type == TriggerUnknown {
		spec.Type = EventTrigger
	}

	return map[string]any{
		"type": string(spec.Type),
	}, nil
}

func SetTriggerType(input json.RawMessage, kind string) (json.RawMessage, error) {
	spec, err := parseTrigger(input)
	if err != nil {
		return nil, err
	}
	spec.Type = TriggerType(kind)
	return encodeInput(spec)
}
