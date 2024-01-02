package datadog

import (
	"github.com/rodrigoprobst/go-plan-management/pkg/configs"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

type StopTracerFunc func()

func InitializeDataDogInstrumentation() StopTracerFunc {

	tracer.Start(
		tracer.WithServiceVersion(configs.ApplicationCfg.AppVersion),
		tracer.WithRuntimeMetrics(),
	)

	return func() {
		tracer.Stop()
	}
}
