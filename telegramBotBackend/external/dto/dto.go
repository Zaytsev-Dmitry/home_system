package dto

type CreateAccountRequestAccountType string

const (
	TG  CreateAccountRequestAccountType = "TG"
	WEB CreateAccountRequestAccountType = "WEB"
)

type AccountDTO struct {
	ID         *int64  `json:"id"`
	FirstName  *string `json:"firstName"`
	LastName   *string `json:"lastName"`
	Username   *string `json:"username"`
	TgUsername *string `json:"telegramUsername"`
	Email      *string `json:"email"`
	TelegramId *int64  `json:"telegramId"`
}
