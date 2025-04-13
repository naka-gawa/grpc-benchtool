package client

import (
	"fmt"
	"log/slog"
	"os"
	"runtime/pprof"
)

type Profiler struct {
	cfg        Config
	cpuProfile *os.File
}

func NewProfiler(cfg Config) *Profiler {
	return &Profiler{cfg: cfg}
}

func (p *Profiler) Start() error {
	if p.cfg.EnableCPUProfile {
		f, err := os.Create(p.cfg.CpuProfilePath)
		if err != nil {
			return fmt.Errorf("failed to create CPU profile file: %w", err)
		}
		p.cpuProfile = f
		if err := pprof.StartCPUProfile(f); err != nil {
			return fmt.Errorf("failed to start CPU profile: %w", err)
		}
		slog.Debug("CPU profiling started", "file", p.cfg.CpuProfilePath)
	} else {
		slog.Debug("CPU profiling is disabled")
	}
	return nil
}

func (p *Profiler) Stop() error {
	if p.cfg.EnableCPUProfile {
		pprof.StopCPUProfile()
		if err := p.cpuProfile.Close(); err != nil {
			return fmt.Errorf("failed to close CPU profile file: %w", err)
		}
		slog.Debug("CPU profiling stopped")
	} else {
		slog.Debug("CPU profiling is disabled")
	}
	return nil
}
