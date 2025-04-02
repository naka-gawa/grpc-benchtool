package datadog

import (
	"encoding/json"
	"time"
)

type metricPayload struct {
	Series []metricSeries `json:"series"`
}

type metricSeries struct {
	Metric string       `json:"metric"`
	Points [][2]float64 `json:"points"`
	Type   string       `json:"type"`
	Tags   []string     `json:"tags,omitempty"`
	Host   string       `json:"host,omitempty"`
}

// buildGaugePayload constructs the JSON payload for a Datadog gauge metric.
func buildGaugePayload(metric string, value float64, tags []string) ([]byte, error) {
	now := float64(time.Now().Unix())
	p := metricPayload{
		Series: []metricSeries{{
			Metric: metric,
			Points: [][2]float64{{now, value}},
			Type:   "gauge",
			Tags:   tags,
		}},
	}
	return json.Marshal(p)
}
