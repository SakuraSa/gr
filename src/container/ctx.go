package container

import (
	"context"

	"github.com/SakuraSa/gr/src/concept"
)

type ContextKey string

const (
	ContainerKey ContextKey = "container"
)

// Get returns the container from the context.
func Get(ctx context.Context) (concept.Container, bool) {
	c, found := ctx.Value(ContainerKey).(concept.Container)
	return c, found
}

// WithContainer returns a new context with the container.
func WithContainer(ctx context.Context) context.Context {
	parent, found := Get(ctx)
	if found && parent != nil {
		return context.WithValue(ctx, ContainerKey, NewTree(parent))
	}
	return context.WithValue(ctx, ContainerKey, NewBase())
}
