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
func (r *PrepareRegistry) prepare(name string, input interface{}) (interface{}, error) {
	fn, ok := r.preparers[name]
	if !ok {
		return nil, fmt.Errorf("preparer %q not found", name)
	}

	return fn(input)
}

// fn func(I) - функция которая выполняется после слоя prepared
func WithPrepared[I any, O any](registry *PrepareRegistry, key string, input I, fn func(I) (O, error)) (O, error) {
	prepared, err := registry.prepare(key, input)
	if err != nil {
		var zero O
		return zero, err
	}
	return fn(prepared.(I))
}

func WithPreparedNoResult[I any](registry *PrepareRegistry, key string, input I, fn func(I) error) error {
	prepared, err := registry.prepare(key, input)
	if err != nil {
		return err
	}
	return fn(prepared.(I))
}
