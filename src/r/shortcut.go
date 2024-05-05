package r

import (
	"context"
	"fmt"

	"github.com/SakuraSa/gr/src/concept"
	"github.com/SakuraSa/gr/src/container"
	"github.com/SakuraSa/gr/src/key"
	"github.com/SakuraSa/gr/src/slot"
)

var (
	ErrContainerNotFound = fmt.Errorf("container not found")
	ErrValueTypeMismatch = fmt.Errorf("value type mismatch")
)

// GetWithKey gets a value with a key from the container in the context.
func GetWithKey[T any](ctx context.Context, k key.Key) (T, error) {
	var t T
	c, found := container.Get(ctx)
	if !found {
		return t, ErrContainerNotFound
	}
	s, err := c.Get(k)
	if err != nil {
		return t, err
	}
	v := s.Value(ctx)
	r, ok := v.Interface().(T)
	if !ok {
		return t, ErrValueTypeMismatch
	}
	return r, err
}

// GetWithName gets a value with a name from the container in the context.
func GetWithName[T any](ctx context.Context, name string) (T, error) {
	return GetWithKey[T](ctx, key.OfNamed[T](name))
}

// Get gets a value from the container in the context, name is default value.
func Get[T any](ctx context.Context) (T, error) {
	return GetWithKey[T](ctx, key.Of[T]())
}

// GetFactoryWithKey gets a factory with a key from the container in the context.
func SetWithKey[T any](ctx context.Context, k key.Key, v T) error {
	c, found := container.Get(ctx)
	if !found {
		return ErrContainerNotFound
	}
	return c.Set(k, slot.NewSingleton(v))
}

// SetWithName sets a value with a name to the container in the context.
func SetWithName[T any](ctx context.Context, name string, v T) error {
	return SetWithKey[T](ctx, key.OfNamed[T](name), v)
}

// Set sets a value to the container in the context, name is default value.
func Set[T any](ctx context.Context, v T) error {
	return SetWithKey[T](ctx, key.Of[T](), v)
}

// SetFactoryWithKey sets a factory with a key to the container in the context.
func SetFactoryWithKey[T any](ctx context.Context, k key.Key, f concept.Factory) error {
	c, found := container.Get(ctx)
	if !found {
		return ErrContainerNotFound
	}
	return c.Set(k, slot.NewFactory(f))
}

// SetFactoryWithName sets a factory with a name to the container in the context.
func SetFactoryWithName[T any](ctx context.Context, name string, f concept.Factory) error {
	return SetFactoryWithKey[T](ctx, key.OfNamed[T](name), f)
}

// SetFactory sets a factory to the container in the context, name is default value.
func SetFactory[T any](ctx context.Context, f concept.Factory) error {
	return SetFactoryWithKey[T](ctx, key.Of[T](), f)
}

// SetFactoryFuncWithKey sets a factory with a key to the container in the context.
func SetFactoryFuncWithKey[T any](ctx context.Context, k key.Key, f concept.FactoryFunc) error {
	c, found := container.Get(ctx)
	if !found {
		return ErrContainerNotFound
	}
	return c.Set(k, slot.NewFactoryWithFunc(f))
}

// SetFactoryFuncWithName sets a factory with a name to the container in the context.
func SetFactoryFuncWithName[T any](ctx context.Context, name string, f concept.FactoryFunc) error {
	return SetFactoryFuncWithKey[T](ctx, key.OfNamed[T](name), f)
}

// SetFactoryFunc sets a factory to the container in the context, name is default value.
func SetFactoryFunc[T any](ctx context.Context, f concept.FactoryFunc) error {
	return SetFactoryFuncWithKey[T](ctx, key.Of[T](), f)
}

// New creates a new context with a container.
func New(ctx context.Context) context.Context {
	return container.WithContainer(ctx)
}
