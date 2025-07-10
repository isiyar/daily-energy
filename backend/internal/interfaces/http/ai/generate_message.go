package ai

import (
	"encoding/json"
)

func GenerateMessage(systemPrompt SystemPrompt, userPrompt string) ([]byte, error) {
	requestBody := ChatRequest{
		Model: "openrouter/cypher-alpha:free",
		Messages: []Message{
			{
				Role:    "system",
				Content: string(systemPrompt),
			},
			{
				Role:    "user",
				Content: userPrompt,
			},
		},
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}
