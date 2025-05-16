package prepare

import (
	"fmt"
)

type PrepareFunc func(input interface{}) (interface{}, error)

type PrepareRegistry struct {
	preparers map[string]PrepareFunc
}

func NewPrepareRegistry() *PrepareRegistry {
	return &PrepareRegistry{
		preparers: make(map[string]PrepareFunc),
	}
}

func RegisterPreparer(r *PrepareRegistry, name string, p Preparer) {
	r.preparers[name] = p.Prepare
}

// Prepare вызывает preparer и делает кастинг уже на стороне вызывающего
func (r *PrepareRegistry) Prepare(name string, input interface{}) (interface{}, error) {
	fn, ok := r.preparers[name]
	if !ok {
		return nil, fmt.Errorf("preparer %q not found", name)
	}

	return fn(input)
}
