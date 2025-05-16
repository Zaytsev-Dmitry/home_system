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

func (r *PrepareRegistry) register(name string, fn PrepareFunc) {
	r.preparers[name] = fn
}

func RegisterPreparer[In any, Out any](r *PrepareRegistry, name string, p Preparer[In, Out]) {
	r.register(name, func(input interface{}) (interface{}, error) {
		inTyped, ok := input.(In)
		if !ok {
			return nil, fmt.Errorf("invalid input type for preparer %q", name)
		}
		return p.Prepare(inTyped)
	})
}

// Prepare вызывает preparer и делает кастинг уже на стороне вызывающего
func (r *PrepareRegistry) Prepare(name string, input interface{}) (interface{}, error) {
	fn, ok := r.preparers[name]
	if !ok {
		return nil, fmt.Errorf("preparer %q not found", name)
	}

	return fn(input)
}
