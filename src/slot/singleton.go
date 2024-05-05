package slot

import (
	"context"
	"reflect"

	"github.com/SakuraSa/gr/src/concept"
)

var (
	_ concept.Slot = (*Singleton)(nil)
)

// Singleton slot is a slot that contains a single value.
type Singleton struct {
	value reflect.Value
}

func NewSingleton(value any) *Singleton {
	return &Singleton{
		value: reflect.ValueOf(value),
	}
}

func NewSingletonWithValue(value reflect.Value) *Singleton {
	return &Singleton{
		value: value,
	}
}

func (s Singleton) Value(_ context.Context) reflect.Value {
	return s.value
}
