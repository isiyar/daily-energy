package ai

import (
	"encoding/json"
	"errors"
)

func Deserialization(bodyBytes []byte, dto *APIResponse) error {
	if err := json.Unmarshal(bodyBytes, &dto); err != nil {
		return errors.New("Failed to parse AI response")
	}

	if len(dto.Choices) == 0 {
		return errors.New("No choices in AI response")
	}
	return nil
}
