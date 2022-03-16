package prober

import (
	"context"
	"time"

	"github.com/m-yosefpor/httpmon/internal/model"
	"github.com/m-yosefpor/httpmon/pkg/metrics"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type Prober struct {
	Logger    *zap.Logger
	Metrics   *metrics.Metrics
	Tracer    trace.Tracer
	ProberCfg Config
}

type ProbeTarget struct {
	ticker *time.Ticker
}

func New(pc Config, l *zap.Logger, mc *metrics.Metrics, t trace.Tracer) *Prober {
	return &Prober{
		Logger:    l,
		Metrics:   mc,
		Tracer:    t,
		ProberCfg: pc,
	}
}

func (p *Prober) NewEndpoint(ctx context.Context, ep model.Endpoint) {
	p.Logger.Info("new endpoint")
	ticker := time.NewTicker(ep.Interval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			p.Logger.Info("Context is done!")
			return
		case <-ticker.C:
			go (func() {
				err := p.sendRequest(ctx, ep.URL)
				if err != nil {
					u.IncFailure(ep)
					u.DecRemain(ep)
				}
				u.IncSuccess(ep)
				u.ZeroRemain(ep) // maybe we should keep last state in memory to avoid updating it unnecessrily
			})()
		}
	}
}

func (p *Prober) Start() {
}
