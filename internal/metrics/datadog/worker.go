package datadog

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"sort"
	"strings"
	"sync"
	"time"
)

type metricKey string

type Metric struct {
	Name  string
	Value float64
	Tags  []string
}

type BufferedClient struct {
	client   *DatadogClient
	buffer   map[metricKey]*Metric
	mu       sync.Mutex
	interval time.Duration
	ctx      context.Context
	cancel   context.CancelFunc
	ticker   *time.Ticker
}

func NewBufferedClient(interval time.Duration) (*BufferedClient, error) {
	dc, err := NewDatadogClient()
	if err != nil {
		return nil, fmt.Errorf("failed to init datadog client: %w", err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	return &BufferedClient{
		client:   dc,
		buffer:   make(map[metricKey]*Metric),
		interval: interval,
		ctx:      ctx,
		cancel:   cancel,
		ticker:   time.NewTicker(interval),
	}, nil
}

func (b *BufferedClient) Add(name string, value float64, tags []string) {
	key := newMetricKey(name, tags)

	b.mu.Lock()
	defer b.mu.Unlock()

	if existing, ok := b.buffer[key]; ok {
		existing.Value += value
	} else {
		b.buffer[key] = &Metric{
			Name:  name,
			Value: value,
			Tags:  tags,
		}
	}
}

func (b *BufferedClient) Start() {
	go func() {
		for {
			select {
			case <-b.ticker.C:
				b.flush()
			case <-b.ctx.Done():
				b.flush()
				return
			}
		}
	}()
}

func (b *BufferedClient) Stop() {
	b.ticker.Stop()
	b.cancel()
}

func (b *BufferedClient) flush() {
	b.mu.Lock()
	defer b.mu.Unlock()

	if len(b.buffer) == 0 {
		return
	}

	now := float64(time.Now().Unix())
	series := make([]metricSeries, 0, len(b.buffer))
	for _, m := range b.buffer {
		series = append(series, metricSeries{
			Metric: m.Name,
			Points: [][2]float64{{now, m.Value}},
			Type:   "count",
			Tags:   m.Tags,
		})
	}

	payload, err := json.Marshal(metricPayload{Series: series})
	if err != nil {
		slog.Error("failed to marshal datadog payload", "error", err)
		return
	}

	if err := sendToDatadog(b.client.client, b.client.apiKey, b.client.appKey, payload); err != nil {
		slog.Error("failed to send datadog metrics", "error", err)
		return
	}

	slog.Debug("flushed datadog metrics", "count", len(series))
	b.buffer = make(map[metricKey]*Metric)
}

func newMetricKey(name string, tags []string) metricKey {
	sort.Strings(tags)
	return metricKey(fmt.Sprintf("%s|%s", name, strings.Join(tags, ",")))
}
