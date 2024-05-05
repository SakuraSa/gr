package util

import "reflect"

// TypeOf returns the reflect.Type of the given type.
func TypeOf[T any]() reflect.Type {
	var v T
	return reflect.TypeOf(v)
}
