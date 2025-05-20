package delegate

import "expensia/internal/app/prepare"

// Обёртки для usecase-функций с подготовкой
func MakeDelegateWithResult[Req any, Resp any](
	registry *prepare.PrepareRegistry,
	key string,
	usecase func(Req) (Resp, error),
) func(Req) (Resp, error) {
	if key == "" || registry == nil {
		return usecase
	}
	return func(req Req) (Resp, error) {
		prepared, err := prepare.Prepare[Req, Req](registry, key, req)
		if err != nil {
			var zero Resp
			return zero, err
		}
		return usecase(prepared)
	}
}

func MakeDelegateNoResult[Req any](
	registry *prepare.PrepareRegistry,
	key string,
	usecase func(Req) error,
) func(Req) error {
	if key == "" || registry == nil {
		return usecase
	}
	return func(req Req) error {
		prepared, err := prepare.Prepare[Req, Req](registry, key, req)
		if err != nil {
			return err
		}
		return usecase(prepared)
	}
}
