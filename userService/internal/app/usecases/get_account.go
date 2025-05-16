package usecases

type GetAccountUCase interface {
	//Get(telegramId int64) (domain.Account, error)
	GetAccountIdByTgId(tgId int64) (accId int64)
}
