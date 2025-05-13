package usecases

import "userService/internal/app/domain"

type GetProfileUCase interface {
	GetByTGId(accId int64) (*domain.Profile, error)
}
