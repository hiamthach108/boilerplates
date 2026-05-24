# Dreon Backend Service Boilerplate

Go backend service boilerplate matching the current `dreon-auth` and `dreon-notification` style.

## Included

- Uber Fx application wiring in `main.go`
- Echo HTTP server under `presentation/http`
- gRPC lifecycle shell under `presentation/grpc`
- Config loading from `.env` plus environment overrides
- `dreon-sdk@v0.0.1` for shared logger, app errors, and generated gRPC contracts/clients
- Redis cache, Postgres/Gorm database client, and validator helpers
- Example layered resource across `aggregate`, `model`, `repository`, `service`, and `handler`

## Start A New Service

1. Copy this directory to the new service repository.
2. Initialize names:

```sh
scripts/init-service.sh github.com/hiamthach108/<new-service> <new-service>
```

3. Update `APP_NAME`, `POSTGRES_DBNAME`, ports, and README.
4. Replace the example resource with service-specific aggregates, models, repositories, services, and handlers.
5. Import shared SDK packages instead of copying service-neutral helpers:

```go
import (
	"github.com/hiamthach108/dreon-sdk/errorx"
	sdklogger "github.com/hiamthach108/dreon-sdk/logger"
	authclient "github.com/hiamthach108/dreon-sdk/clients/auth"
	noticlient "github.com/hiamthach108/dreon-sdk/clients/notification"
)
```

6. Run:

```sh
go mod tidy
go test ./...
```

## Local Run

```sh
docker compose up -d
make run
```
