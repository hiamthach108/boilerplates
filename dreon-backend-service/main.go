package main

import (
	"github.com/hiamthach108/dreon-backend-service/config"
	"github.com/hiamthach108/dreon-backend-service/internal/repository"
	"github.com/hiamthach108/dreon-backend-service/internal/service"
	"github.com/hiamthach108/dreon-backend-service/pkg/cache"
	"github.com/hiamthach108/dreon-backend-service/pkg/database"
	"github.com/hiamthach108/dreon-backend-service/pkg/logger"
	grpcserver "github.com/hiamthach108/dreon-backend-service/presentation/grpc"
	"github.com/hiamthach108/dreon-backend-service/presentation/http"
	"github.com/hiamthach108/dreon-backend-service/presentation/http/handler"
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
			logger.NewLogger,
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
