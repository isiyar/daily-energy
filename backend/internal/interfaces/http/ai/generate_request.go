package ai

import (
	"bytes"
	"fmt"
	"github.com/isiyar/daily-energy/backend/config"
	"net/http"
)

func GenerateRequest(config config.Config, jsonData []byte) (*http.Request, error) {
	req, err := http.NewRequest("POST", config.ApiPath, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.ApiKey))
	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("HTTP-Referer", "<YOUR_SITE_URL>") // Optional
	//req.Header.Set("X-Title", "<YOUR_SITE_NAME>")     // Optional

	return req, nil
}
