package server

import (
	"net"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/m-yosefpor/httpmon/pkg/bind"
	"github.com/m-yosefpor/httpmon/pkg/metrics"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type Server struct {
	Logger  *zap.Logger
	Metrics *metrics.Metrics
	Tracer  trace.Tracer
	Bind    bind.Config
}

func New(b bind.Config, l *zap.Logger, mc *metrics.Metrics, t trace.Tracer) *Server {
	return &Server{
		Logger:  l,
		Metrics: mc,
		Tracer:  t,
		Bind:    b,
	}
}

func (s *Server) Serve() {
	s.Logger.Info("starting")
	app := fiber.New()
	app.Get("/hi", Hi)
	addr := net.JoinHostPort(s.Bind.Host, strconv.Itoa(s.Bind.Port))
	if err := app.Listen(addr); err != nil {
		s.Logger.Panic("cannot start the server")
	}
}

func (s *Server) Shutdown() {
	s.Logger.Info("shutting down")
	time.Sleep(1 * time.Second)
}

func Hi(c *fiber.Ctx) error {
	return c.SendString("hello")
}
