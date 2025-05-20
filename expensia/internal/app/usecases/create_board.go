package usecases

import (
	"expensia/internal/app/domain"
)

type CreateBoardUCase interface {
	CreateAndReturnBoard(req CreateBoardInput) (*domain.Board, error)
}
type CreateBoardInput struct {
	TgUserId int64
	OwnerId  int64
	Name     string
	Currency string
}
