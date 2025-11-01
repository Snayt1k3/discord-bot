package main

import (
	"context"
	"fmt"
	"interaction-service/config"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
	// "interaction-service/internal/adapters/repos"
	"interaction-service/internal/models"
	"interaction-service/internal/server"
	pb "interaction-service/proto"
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

	srvMetrics := grpcprom.NewServerMetrics(
		grpcprom.WithServerHandlingTimeHistogram(
			grpcprom.WithHistogramBuckets([]float64{0.001, 0.01, 0.1, 0.3, 0.6, 1, 3, 6, 9, 20, 30, 60, 90, 120}),
		),
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

	pb.RegisterInteractionServiceServer(grpcServer, server.NewUserServer(db))
	srvMetrics.InitializeMetrics(grpcServer)

	g := &run.Group{}
	addGrpcServer(g, grpcServer, cfg.GrpcPort)
	addPrometheusServer(g, reg, 8081)

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

func addPrometheusServer(g *run.Group, reg *prometheus.Registry, port int) {
	httpSrv := &http.Server{Addr: fmt.Sprintf(":%d", port)}

	g.Add(func() error {
		mux := http.NewServeMux()
		mux.Handle("/metrics", promhttp.HandlerFor(
			reg,
			promhttp.HandlerOpts{EnableOpenMetrics: true},
		))
		httpSrv.Handler = mux
		slog.Info("starting HTTP server", "addr", httpSrv.Addr)
		return httpSrv.ListenAndServe()
	}, func(error) {
		slog.Info("stopping HTTP server")
		if err := httpSrv.Close(); err != nil {
			slog.Error("failed to stop web server", "err", err)
		}
	})
}

func addGrpcServer(g *run.Group, grpcServer *grpc.Server, port string) {
	g.Add(func() error {
		l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
		if err != nil {
			return err
		}
		slog.Info("starting gRPC server", "addr", l.Addr().String())
		return grpcServer.Serve(l)
	}, func(err error) {
		slog.Info("stopping gRPC server")
		grpcServer.GracefulStop()
	})
}
