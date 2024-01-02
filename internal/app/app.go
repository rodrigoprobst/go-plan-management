package app

import (
	"context"
	"github.com/rodrigoprobst/go-plan-management/internal/resolver"
)

var App *Application

type Application struct {
	AppCtx   context.Context
	Resolver resolver.Resolver
}

func NewApplication(ctx context.Context, resolver resolver.Resolver) {
	App = &Application{
		AppCtx:   ctx,
		Resolver: resolver,
	}
}
