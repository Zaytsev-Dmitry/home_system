package delegate

import "expensia/internal/app/prepare"

// For usecases with result
func MakeDelegateWithResult[Req any, Resp any](
	registry *prepare.PrepareRegistry,
	key string,
	usecase func(Req) (Resp, error),
) func(Req) (Resp, error) {
	if key == "" || registry == nil {
		return usecase
	} else {
		return func(req Req) (Resp, error) {
			return prepare.WithPrepared(registry, key, req, usecase)
		}
	}
}

// For usecases without result
func MakeDelegateNoResult[Req any](
	registry *prepare.PrepareRegistry,
	key string,
	usecase func(Req) error,
) func(Req) error {
	if key == "" || registry == nil {
		return usecase
	} else {
		return func(req Req) error {
			return prepare.WithPreparedNoResult(registry, key, req, usecase)
		}
	}
}
