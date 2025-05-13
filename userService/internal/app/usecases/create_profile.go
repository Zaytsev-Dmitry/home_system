package usecases

import "userService/internal/app/domain"

type CreateProfileUCase interface {
	Create(acc domain.Account) error
}
