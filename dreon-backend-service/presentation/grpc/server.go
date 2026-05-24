package grpc

import (
	"context"
	"net"
	"time"

	"github.com/hiamthach108/dreon-backend-service/config"
	"github.com/hiamthach108/dreon-backend-service/pkg/logger"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	server *grpc.Server
	config *config.AppConfig
	logger logger.ILogger
}

func NewGRPCServer(
	cfg *config.AppConfig,
	logger logger.ILogger,
) *GRPCServer {
	s := grpc.NewServer()
	reflection.Register(s)
	return &GRPCServer{
		server: s,
		config: cfg,
		logger: logger,
	}
}

func RegisterHooks(lc fx.Lifecycle, srv *GRPCServer) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			port := srv.config.Server.GRPCPort
			if port == "" {
				port = "9090"
			}
			addr := net.JoinHostPort(srv.config.Server.Host, port)
			lis, err := net.Listen("tcp", addr)
			if err != nil {
				return err
			}
			srv.logger.Info("Starting gRPC server", "addr", addr)
			go func() {
				if err := srv.server.Serve(lis); err != nil {
					srv.logger.Fatal("gRPC server failed", "error", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			srv.logger.Info("Shutting down gRPC server...")
			stopped := make(chan struct{})
			go func() {
				srv.server.GracefulStop()
				close(stopped)
			}()
			select {
			case <-ctx.Done():
				srv.server.Stop()
			case <-stopped:
			case <-time.After(5 * time.Second):
				srv.server.Stop()
			}
			return nil
		},
	})
}
