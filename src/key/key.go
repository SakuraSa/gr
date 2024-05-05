package key

import (
	"reflect"
	"strings"

	"github.com/SakuraSa/gr/src/util"
)

const (
	DefaultName = "default"
)

// Key is a type that represents a key for a value in a container.
type Key struct {
	Type reflect.Type
	Name string
}

// String returns a string representation of the key.
func (k Key) String() string {
	var sb strings.Builder
	sb.WriteString(k.Type.Name())
	if k.Name != DefaultName {
		sb.WriteString("#")
		sb.WriteString(k.Name)

	}
	return sb.String()
}

// Of returns a new key with the type of the given value.
func Of[T any]() Key {
	return Key{
		Type: util.TypeOf[T](),
		Name: DefaultName,
	}
}

// OfNamed returns a new key with the type of the given value and the given name.
func OfNamed[T any](name string) Key {
	return Key{
		Type: util.TypeOf[T](),
		Name: name,
	}
}
