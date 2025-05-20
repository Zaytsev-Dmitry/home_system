package prepare

// Универсальный интерфейс препарера
type Preparer[I any, O any] interface {
	Prepare(input I) (O, error)
}
