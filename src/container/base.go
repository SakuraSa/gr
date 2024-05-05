package container

import (
	"github.com/SakuraSa/gr/src/concept"
	"github.com/SakuraSa/gr/src/key"
)

var (
	_ concept.Container = (*Base)(nil)
)

type Base struct {
	valueMap map[key.Key]concept.Slot
}

func NewBase() *Base {
	return &Base{
		valueMap: make(map[key.Key]concept.Slot),
	}
}

func (b *Base) Get(k key.Key) (concept.Slot, error) {
	if v, ok := b.valueMap[k]; ok {
		return v, nil
	}
	return nil, ErrKeyNotFound
}

func (b *Base) Set(k key.Key, v concept.Slot) error {
	if _, ok := b.valueMap[k]; ok {
		return ErrKeyAlreadyExists
	}
	b.valueMap[k] = v
	return nil
}
