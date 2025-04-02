package server

import "time"

type Config struct {
	Host                string
	Port                int
	Timeout             time.Duration
	ServerID            string
	DatadogCustomMetric bool
}
