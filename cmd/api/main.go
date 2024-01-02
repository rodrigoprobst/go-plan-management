package main

import (
	"context"

	"github.com/rodrigoprobst/go-plan-management/internal/app"
	"github.com/rodrigoprobst/go-plan-management/internal/resolver"
	"github.com/rodrigoprobst/go-plan-management/pkg/configs"
	"github.com/rodrigoprobst/go-plan-management/pkg/datadog"
	"github.com/rodrigoprobst/go-plan-management/pkg/logger"
	"github.com/rodrigoprobst/go-plan-management/pkg/postgres"
	"github.com/rodrigoprobst/go-plan-management/pkg/server/http_server"
	"github.com/rodrigoprobst/go-plan-management/pkg/validation"

	"go.uber.org/zap"
)

func main() {
	configs.InitializeConfigs()

	logr, err := logger.NewLogger(configs.ApplicationCfg.AppName, configs.ApplicationCfg.Env)
	if err != nil {
		panic(err)
	}
	if configs.ApplicationCfg.Env == configs.Production {
		shutdownTrace := datadog.InitializeDataDogInstrumentation()
		defer shutdownTrace()
	}

	logr = logr.With(zap.String("version", configs.ApplicationCfg.AppVersion))

	validation.InitializeValidatorConfigs()

	appContext := context.Background()
	postgresDatabase := postgres.GetDatabase(logr)
	reslv := resolver.NewResolver(resolver.WithPostgresDatabase(postgresDatabase))
	app.NewApplication(appContext, reslv)

	http_server.RunHttpServer(logr)
}
