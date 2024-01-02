package routes

import (
	"github.com/rodrigoprobst/go-plan-management/internal/handlers"
	"github.com/rodrigoprobst/go-plan-management/pkg/server/http_server/http_middlewares"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Routes(engine *gin.Engine, logger *zap.Logger) *gin.Engine {

	engine.GET("/health-check", handlers.HealthcheckHandler)
	engine.MaxMultipartMemory = 1 << 20

	api := engine.Group("/", http_middlewares.LoggerMiddleware(logger))
	{
		api.GET("/liveness", handlers.HealthcheckHandler)
		api.GET("/readiness", handlers.HealthcheckHandler)
	}

	return engine
}
