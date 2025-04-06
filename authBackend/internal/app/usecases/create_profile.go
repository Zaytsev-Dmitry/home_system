package usecases

import "authServer/internal/app/domain"

type CreateProfileUCase interface {
	Create(acc domain.Account) error
}
