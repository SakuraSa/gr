package concept

import (
	"context"
	"reflect"
)

// Slot .
type Slot interface {
	Value(context.Context) reflect.Value
}
