package datadog

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type DatadogClient struct {
	apiKey string
	appKey string
	client *http.Client
}

var ErrMissingCredentials = fmt.Errorf("missing DATADOG_API_KEY or DATADOG_APP_KEY")

func NewDatadogClient() (*DatadogClient, error) {
	apiKey := os.Getenv("DATADOG_API_KEY")
	appKey := os.Getenv("DATADOG_APP_KEY")
	if apiKey == "" || appKey == "" {
		slog.Error("missing required Datadog credentials: DATADOG_API_KEY or DATADOG_APP_KEY")
		return nil, ErrMissingCredentials
	}

	return &DatadogClient{
		apiKey: apiKey,
		appKey: appKey,
		client: &http.Client{Timeout: 5 * time.Second},
	}, nil
}

func (d *DatadogClient) SendGauge(metric string, value float64, tags []string) error {
	payload, err := buildGaugePayload(metric, value, tags)
	if err != nil {
		slog.Error("failed to build payload", slog.Any("error", err))
		return err
	}
	return sendToDatadog(d.client, d.apiKey, d.appKey, payload)
}
