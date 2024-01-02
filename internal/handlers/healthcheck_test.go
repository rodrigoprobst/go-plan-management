package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/rodrigoprobst/go-plan-management/pkg/configs"
	"github.com/rodrigoprobst/go-plan-management/pkg/helpers/test_helpers/gin_test_functions"
	"github.com/rodrigoprobst/go-plan-management/pkg/helpers/test_helpers/test_functions"

	"github.com/gin-gonic/gin"
)

func TestHealthcheckHandler(t *testing.T) {
	expectedResponse := "{\"status\":\"available\",\"system_info\":{\"environment\":\"\",\"version\":\"v.test\"}}"

	configs.InitializeConfigs()
	configs.ApplicationCfg.AppVersion = "v.test"
	w := httptest.NewRecorder()

	gin.SetMode(gin.ReleaseMode)
	var params []gin.Param
	u := url.Values{}
	engine, ctx := gin_test_functions.BuildGinTestEngine(w)

	engine.GET("/ping", HealthcheckHandler)
	gin_test_functions.MockJsonGet(ctx, params, u)
	req, _ := http.NewRequestWithContext(ctx, "GET", "/ping", nil)
	engine.ServeHTTP(w, req)
	responseData, _ := io.ReadAll(w.Body)

	test_functions.IsObjectEqual(t, string(responseData), expectedResponse)
}
