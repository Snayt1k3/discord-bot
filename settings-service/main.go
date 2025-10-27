package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
	"settings-service/config"
	"settings-service/internal/adapters/repos"
	"settings-service/internal/models"
	"settings-service/internal/server"
	pb "settings-service/proto"
	"syscall"

	grpcprom "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/oklog/run"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"google.golang.org/grpc"
)

func main() {
	initLogging()
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	models.Migrate(db)
	repositories := repos.NewRepos(db)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	srvMetrics := grpcprom.NewServerMetrics(
		grpcprom.WithServerHandlingTimeHistogram(
			grpcprom.WithHistogramBuckets([]float64{0.001, 0.01, 0.1, 0.3, 0.6, 1, 3, 6, 9, 20, 30, 60, 90, 120}),
		),
		// Add tenant_name as a context label. This server option is necessary
		// to initialize the metrics with the labels that will be provided
		// dynamically from the context. This should be used in tandem with
		// WithLabelsFromContext in the interceptor options.
		// grpcprom.WithContextLabels("tenant_name"),
	)
	reg := prometheus.NewRegistry()
	reg.MustRegister(
		srvMetrics,
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))

	grpcServer := grpc.NewServer(
		grpc.ChainStreamInterceptor(srvMetrics.StreamServerInterceptor()),
		grpc.UnaryInterceptor(srvMetrics.UnaryServerInterceptor()),
	)

	pb.RegisterAutoModServiceServer(grpcServer, &server.AutomodeServer{
		Repo: repositories.AutoMode, GuildRepo: repositories.Settings,
	})
	pb.RegisterLogServiceServer(grpcServer, &server.LogServer{
		Repo: repositories.Log,
	})
	pb.RegisterRolesServiceServer(grpcServer, &server.RolesReactionServer{
		Repo: repositories.ReactionRoles, GuildRepo: repositories.Settings,
	})
	pb.RegisterSettingsServiceServer(grpcServer, &server.GuildServer{
		Repo: repositories.Settings,
	})
	pb.RegisterWelcomeServiceServer(grpcServer, &server.WelcomeServer{
		Repo: repositories.Welcome, GuildRepo: repositories.Settings,
	})

	srvMetrics.InitializeMetrics(grpcServer)

	g := &run.Group{}
	g.Add(func() error {
		l, err := net.Listen("tcp", fmt.Sprintf(":%v", cfg.GrpcPort))
		if err != nil {
			return err
		}
		slog.Info("starting gRPC server", "addr", l.Addr().String())
		return grpcServer.Serve(l)
	}, func(err error) {
		grpcServer.GracefulStop()
		grpcServer.Stop()
	})

	httpSrv := &http.Server{Addr: ":8081"}
	g.Add(func() error {
		m := http.NewServeMux()
		// Create HTTP handler for Prometheus metrics.
		m.Handle("/metrics", promhttp.HandlerFor(
			reg,
			promhttp.HandlerOpts{
				EnableOpenMetrics: true,
			},
		))
		httpSrv.Handler = m
		slog.Info("starting HTTP server", "addr", httpSrv.Addr)
		return httpSrv.ListenAndServe()
	}, func(error) {
		if err := httpSrv.Close(); err != nil {
			slog.Error("failed to stop web server", "err", err)
		}
	})

	g.Add(run.SignalHandler(context.Background(), syscall.SIGINT, syscall.SIGTERM))

	if err := g.Run(); err != nil {
		slog.Error("program interrupted", "err", err)
		os.Exit(1)
	}
}

func initLogging() {
	opts := &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: true,
	}
	// Вывод логов в консоль
	logger := slog.New(slog.NewTextHandler(os.Stdout, opts))
	slog.SetDefault(logger)
	slog.Info("Logger initialized")
}
