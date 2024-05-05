package concept

import (
	"context"
	"reflect"
)

// FactoryFunc is a function that creates a new instance of a type.
type FactoryFunc func(context.Context) reflect.Value

// Factory is a type that creates a new instance of a type.
type Factory interface {
	NewInstance(context.Context) reflect.Value
}
