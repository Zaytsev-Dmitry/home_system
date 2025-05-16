package repository

type CreateBoardUCaseIn struct {
	TgUserId int64
	OwnerId  int64
	Name     string
	Currency string
}
