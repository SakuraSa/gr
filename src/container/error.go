package container

import "fmt"

var (
	ErrKeyNotFound      = fmt.Errorf("key not found")
	ErrKeyAlreadyExists = fmt.Errorf("key already exists")
)
