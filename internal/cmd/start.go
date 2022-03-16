package cmd

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/m-yosefpor/httpmon/pkg/logger"
	"github.com/m-yosefpor/httpmon/pkg/metrics"
	"github.com/m-yosefpor/httpmon/pkg/profiler"
	"github.com/m-yosefpor/httpmon/pkg/tracer"

	"github.com/m-yosefpor/httpmon/internal/config"
	"github.com/m-yosefpor/httpmon/internal/db"
	"github.com/m-yosefpor/httpmon/internal/http/handler"
	"github.com/m-yosefpor/httpmon/internal/prober"
	"github.com/m-yosefpor/httpmon/internal/server"
	"github.com/m-yosefpor/httpmon/internal/store"
)

func newStartCmd() *cobra.Command {

	startCmd := &cobra.Command{
		Use:   "start",
		Short: "start the webserver",
		Run: func(cmd *cobra.Command, args []string) {
			start()
		},
	}
	return startCmd
}

func start() {

	// initialize configuration
	c := config.Load()

	// initialize logger
	l := logger.New(c.Logger)

	// initialize metrics
	mc, err := metrics.New(c.Metrics)
	if err != nil {
		l.Fatal("failed to create metrics: %w", zap.Error(err))
	}

	// initialize tracer
	t := tracer.New(c.Tracer)

	// initialize profiler
	profiler.Start(c.Profiler)

	// connect to db
	db, err := db.New(c.Database)
	if err != nil {
		log.Fatalf("database connection failed: %s", err)
	}

	hs := handler.User{
		Store: store.NewMongoDBStore(db),
	}

	// initialize server
	app := fiber.New()
	app.Post("/create_user", hs.CreateUser)
	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))
	g := app.Group("/")
	hs.Register(g)

	if err := app.Listen(":1373"); err != nil {
		log.Println("cannot start the server")
	}

	s := server.New(c.Bind, l, mc, t)

	go s.Serve()

	// initilize prober
	p := prober.New(c.Prober, l, mc, t)
	go p.Start()

	// signal handling
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	// db cancel
	// prober cancel
	// server shutdown
	s.Shutdown()
}
