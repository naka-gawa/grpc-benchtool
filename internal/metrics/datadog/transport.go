package datadog

import (
	"bytes"
	"fmt"
	"log/slog"
	"net/http"
)

// sendToDatadog sends the payload to the Datadog metrics API endpoint.
func sendToDatadog(client *http.Client, apiKey, appKey string, payload []byte) error {
	req, err := http.NewRequest("POST", "https://api.datadoghq.com/api/v1/series", bytes.NewBuffer(payload))
	if err != nil {
		slog.Error("failed to create request", slog.Any("error", err))
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("DD-API-KEY", apiKey)
	req.Header.Set("DD-APPLICATION-KEY", appKey)

	resp, err := client.Do(req)
	if err != nil {
		slog.Error("failed to send request", slog.Any("error", err))
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		slog.Error("datadog returned error", slog.Int("status_code", resp.StatusCode))
		return fmt.Errorf("datadog returned status: %s", resp.Status)
	}

	slog.Debug("datadog metric sent successfully")
	return nil
}
