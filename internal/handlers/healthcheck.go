package handlers

import (
	"net/http"

	"github.com/rodrigoprobst/go-plan-management/pkg/configs"

	"github.com/gin-gonic/gin"
)

func HealthcheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "available",
		"system_info": map[string]string{
			"environment": configs.ApplicationCfg.Env,
			"version":     configs.ApplicationCfg.AppVersion,
		},
	})
}
