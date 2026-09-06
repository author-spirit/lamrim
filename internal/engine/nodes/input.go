package nodes

import "encoding/json"

func decodeInput(input json.RawMessage, dst any) error {
	if len(input) == 0 {
		return nil
	}
	return json.Unmarshal(input, dst)
}

func encodeInput(src any) (json.RawMessage, error) {
	return json.Marshal(src)
}
