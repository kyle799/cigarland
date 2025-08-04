package main

import (
	"encoding/json"
	"fmt"
)

func UnmarshalJsonPayload(payload []byte, target any) (unmarshalError error) {
	unmarshalError = json.Unmarshal(payload, target)
	if unmarshalError != nil {
		return fmt.Errorf("Error unmarshaling json into target: %w", unmarshalError)
	}
	return nil
}
