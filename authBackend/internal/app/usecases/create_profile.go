package usecases

import "authBackend/internal/app/domain"

type CreateProfileUCase interface {
	Create(acc domain.Account) error
}
