package config

import (
	"github.com/m-yosefpor/httpmon/pkg/bind"
	"github.com/m-yosefpor/httpmon/pkg/koanfconf"
	"github.com/m-yosefpor/httpmon/pkg/logger"
	"github.com/m-yosefpor/httpmon/pkg/metrics"
	"github.com/m-yosefpor/httpmon/pkg/profiler"
	"github.com/m-yosefpor/httpmon/pkg/tracer"

	"github.com/m-yosefpor/httpmon/internal/db"
	"github.com/m-yosefpor/httpmon/internal/prober"
)

const (
	envPrefix      = "httpmon_"
	configFileName = "config.yml"
)

type Config struct {
	Logger   logger.Config   `koanf:"logger"`
	Metrics  metrics.Config  `koanf:"metrics"`
	Tracer   tracer.Config   `koanf:"tracer"`
	Profiler profiler.Config `koanf:"profiler"`
	Bind     bind.Config     `koanf:"bind"`
	Database db.Config       `koanf:"database"`
	Prober   prober.Config   `koanf:"prober"`
}

func Load() Config {
	var instance Config
	koanfconf.Load(configFileName, envPrefix, Default(), &instance)
	return instance
}
