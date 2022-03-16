package config

import (
	"time"

	"github.com/m-yosefpor/httpmon/pkg/bind"
	"github.com/m-yosefpor/httpmon/pkg/logger"
	"github.com/m-yosefpor/httpmon/pkg/metrics"
	"github.com/m-yosefpor/httpmon/pkg/profiler"
	"github.com/m-yosefpor/httpmon/pkg/tracer"

	"github.com/m-yosefpor/httpmon/internal/db"
	"github.com/m-yosefpor/httpmon/internal/prober"
)

// nolint: gomnd, funlen
func Default() Config {
	return Config{
		Logger: logger.Config{
			Level: "warn",
			Syslog: logger.Syslog{
				Enabled: false,
				Network: "",
				Address: "",
				Tag:     "",
			},
		},
		Metrics: metrics.Config{},
		Tracer: tracer.Config{
			Enabled: false,
			Agent: tracer.Agent{
				Host: "127.0.0.1",
				Port: "6831",
			},
			Ratio: 1.0,
		},
		Profiler: profiler.Config{
			Enabled: false,
			Address: "http://127.0.0.1:4040",
		},
		Bind: bind.Config{
			Host: "0.0.0.0",
			Port: 1378,
		},
		Database: db.Config{
			URL:               "mongodb://127.0.0.1:27017",
			ConnectionTimeout: 10 * time.Second,
			Name:              "httpmon",
		},
		Prober: prober.Config{
			Interval: 10 * time.Second,
		},
	}
}
