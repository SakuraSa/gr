package container

import (
	"github.com/SakuraSa/gr/src/concept"
	"github.com/SakuraSa/gr/src/key"
)

var (
	_ concept.Container = (*Tree)(nil)
)

// Tree container is a container that contains a tree of slots.
type Tree struct {
	self   *Base
	parent concept.Container
}

func NewTree(parent concept.Container) *Tree {
	return &Tree{
		self:   NewBase(),
		parent: parent,
	}
}

func (t *Tree) Get(k key.Key) (concept.Slot, error) {
	if v, err := t.self.Get(k); err == nil {
		return v, nil
	}
	if t.parent == nil {
		return nil, ErrKeyNotFound
	}
	return t.parent.Get(k)
}

func (t *Tree) Set(k key.Key, v concept.Slot) error {
	return t.self.Set(k, v)
}
