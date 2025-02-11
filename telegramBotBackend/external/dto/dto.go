package dto

type CreateAccountRequestAccountType string

const (
	TG  CreateAccountRequestAccountType = "TG"
	WEB CreateAccountRequestAccountType = "WEB"
)

//type AccountResponse struct {
//	Id               *int64  `json:"id,omitempty"`
//	Username         *string `json:"username,omitempty"`
//	FirstName        *string `json:"firstName,omitempty"`
//	LastName         *string `json:"lastName,omitempty"`
//	Email            *string `json:"email,omitempty"`
//	TelegramUserName *string `json:"telegramUserName,omitempty"`
//	TelegramId       *int64  `json:"telegramId,omitempty"`
//}
//
//type CreateAccountRequest struct {
//	AccountType      *CreateAccountRequestAccountType `json:"accountType,omitempty"`
//	Email            *string                          `json:"email,omitempty"`
//	FirstName        *string                          `json:"firstName,omitempty"`
//	LastName         *string                          `json:"lastName,omitempty"`
//	Username         *string                          `json:"username,omitempty"`
//	TelegramUserName *string                          `json:"telegramUserName,omitempty"`
//	TelegramId       *int64                           `json:"telegramId,omitempty"`
//}

type AccountDTO struct {
	ID         int64  `json:"id"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Username   string `json:"username"`
	TgUsername string `json:"telegramUsername"`
	Email      string `json:"email"`
	TelegramId int64  `json:"telegramId"`
}
