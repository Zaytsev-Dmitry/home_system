package echo

import "telegramCLient/internal/storage"

type Option func(e *Echo)

func WithMessageStorage(storage storage.Storage) Option {
	return func(e *Echo) {
		e.messageStorage = storage
	}
}
