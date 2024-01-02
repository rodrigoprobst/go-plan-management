package http_server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/rodrigoprobst/go-plan-management/internal/routes"
	"github.com/rodrigoprobst/go-plan-management/pkg/configs"
	"github.com/rodrigoprobst/go-plan-management/pkg/server/http_server/http_middlewares"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	dd_gin "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
)

func RunHttpServer(logger *zap.Logger) {

	engine := gin.Default()

	engine.Use(http_middlewares.CorsMiddleware())
	engine.Use(dd_gin.Middleware(""))

	srv := http.Server{
		Addr:              fmt.Sprintf(":%s", configs.ApplicationCfg.Port),
		Handler:           routes.Routes(engine, logger),
		ReadHeaderTimeout: 60 * time.Second,
	}

	logger.Info(fmt.Sprintf("Starting server on addr: %s", srv.Addr))
	logger.Fatal("Server unexpected shutdown", zap.Error(srv.ListenAndServe()))
}
