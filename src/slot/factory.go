package slot

import (
	"context"
	"reflect"

	"github.com/SakuraSa/gr/src/concept"
)

// Factory slot is a slot that contains a factory function.
type Factory struct {
	f concept.FactoryFunc
}

func NewFactory(f concept.Factory) *Factory {
	return &Factory{
		f: f.NewInstance,
	}
}

func NewFactoryWithFunc(f concept.FactoryFunc) *Factory {
	return &Factory{
		f: f,
	}
}

func (f Factory) Value(ctx context.Context) reflect.Value {
	return f.f(ctx)
}
