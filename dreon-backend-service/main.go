package main

import (
	"github.com/hiamthach108/dreon-backend-service/config"
	"github.com/hiamthach108/dreon-backend-service/internal/repository"
	"github.com/hiamthach108/dreon-backend-service/internal/service"
	"github.com/hiamthach108/dreon-backend-service/pkg/cache"
	"github.com/hiamthach108/dreon-backend-service/pkg/database"
	grpcserver "github.com/hiamthach108/dreon-backend-service/presentation/grpc"
	"github.com/hiamthach108/dreon-backend-service/presentation/http"
	"github.com/hiamthach108/dreon-backend-service/presentation/http/handler"
	"github.com/hiamthach108/dreon-sdk/logger"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func main() {
	app := fx.New(
		fx.WithLogger(func(appLogger logger.ILogger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: appLogger.GetZapLogger()}
		}),
		fx.Provide(
			// Core
			config.NewAppConfig,
			newAppLogger,
			cache.NewAppCache,
			database.NewDbClient,
			http.NewHttpServer,

			// Repositories
			repository.NewExampleRepository,

			// Services
			service.NewExampleSvc,

			// Handlers
			handler.NewExampleHandler,

			// gRPC server
			grpcserver.NewGRPCServer,
		),
		fx.Invoke(http.RegisterHooks),
		fx.Invoke(grpcserver.RegisterHooks),
	)

	app.Run()
}

func newAppLogger(config *config.AppConfig) (logger.ILogger, error) {
	return logger.NewLogger(logger.Config{
		Service: config.App.Name,
		Level:   config.Logger.Level,
	})
}
