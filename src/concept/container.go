package concept

import "github.com/SakuraSa/gr/src/key"

type Container interface {
	Get(key.Key) (Slot, error)
	Set(key.Key, Slot) error
}
