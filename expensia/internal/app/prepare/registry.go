package prepare

import (
	"fmt"
	"reflect"
	"sync"
)

// Обёртка для функции-препарера
type PrepareFunc[I any, O any] func(input I) (O, error)

type preparerEntry struct {
	inType  reflect.Type
	outType reflect.Type
	fn      any // PrepareFunc[I, O] или Preparer[I, O]
}

// Универсальный потокобезопасный реестр
type PrepareRegistry struct {
	mu        sync.RWMutex
	preparers map[string]*preparerEntry
}

func NewPrepareRegistry() *PrepareRegistry {
	return &PrepareRegistry{
		preparers: make(map[string]*preparerEntry),
	}
}

// Регистрация препарера как структуры
func RegisterPreparer[I any, O any](r *PrepareRegistry, name string, p Preparer[I, O]) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var in I
	var out O
	r.preparers[name] = &preparerEntry{
		inType:  reflect.TypeOf(in),
		outType: reflect.TypeOf(out),
		fn:      p,
	}
}

// Регистрация препарера как функции
func RegisterPrepareFunc[I any, O any](r *PrepareRegistry, name string, fn PrepareFunc[I, O]) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var in I
	var out O
	r.preparers[name] = &preparerEntry{
		inType:  reflect.TypeOf(in),
		outType: reflect.TypeOf(out),
		fn:      fn,
	}
}

// Типобезопасный вызов препарера
func Prepare[I any, O any](r *PrepareRegistry, name string, input I) (O, error) {
	r.mu.RLock()
	entry, ok := r.preparers[name]
	r.mu.RUnlock()

	var zeroO O
	if !ok {
		return zeroO, fmt.Errorf("preparer %q not found", name)
	}

	// Проверка типа входа (для отладки, можно убрать на проде)
	if entry.inType != reflect.TypeOf(input) {
		return zeroO, fmt.Errorf("preparer %q: input type mismatch (want %v, got %v)", name, entry.inType, reflect.TypeOf(input))
	}

	// Вызов по типу
	switch fn := entry.fn.(type) {
	case Preparer[I, O]:
		return fn.Prepare(input)
	case PrepareFunc[I, O]:
		return fn(input)
	default:
		return zeroO, fmt.Errorf("preparer %q: invalid preparer type", name)
	}
}
