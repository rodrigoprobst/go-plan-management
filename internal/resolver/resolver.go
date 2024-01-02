package resolver

import (
	"time"

	"github.com/uptrace/bun"
)

type Interface int

type Resolver interface {
}

type options struct {
	singletonFor map[Interface]bool
	timeout      map[Interface]time.Duration
}

func (o *options) shouldBeSingleton(adapter Interface) bool {
	if should, ok := o.singletonFor[adapter]; ok && should {
		return true
	}
	return false
}

type resolver struct {
	db      *bun.DB
	options options
}

func NewResolver(opts ...ResolverOption) Resolver {
	reslv := &resolver{}
	reslv.options.singletonFor = make(map[Interface]bool)
	reslv.options.timeout = make(map[Interface]time.Duration)
	for _, opt := range opts {
		opt(reslv)
	}
	return reslv
}

type ResolverOption func(*resolver)

func WithSingletonFor(adapters ...Interface) ResolverOption {
	return func(r *resolver) {
		for _, adapter := range adapters {
			r.options.singletonFor[adapter] = true
		}
	}
}

func WithTimeoutFor(adapter Interface, timeout time.Duration) ResolverOption {
	return func(r *resolver) {
		r.options.timeout[adapter] = timeout
	}
}

func WithPostgresDatabase(db *bun.DB) ResolverOption {
	return func(r *resolver) {
		r.db = db
	}
}
