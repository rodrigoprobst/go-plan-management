package http_middlewares

import (
	"net/http"
	"time"

	"github.com/rodrigoprobst/go-plan-management/pkg/configs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	trustedOrigins := configs.ApplicationCfg.CorsTrustedOrigins
	if len(trustedOrigins) == 0 {
		trustedOrigins = []string{"https://localhost", "https://*.localhost"}
	}

	return cors.New(cors.Config{
		AllowOrigins:     trustedOrigins,
		AllowMethods:     []string{http.MethodOptions, http.MethodPut, http.MethodPatch, http.MethodDelete},
		AllowHeaders:     []string{"Authorization", "Content-Type", "Cache-Control", "X-XSRF-TOKEN"},
		AllowCredentials: true,
		AllowWildcard:    true,
		MaxAge:           60 * time.Second,
	})
}
