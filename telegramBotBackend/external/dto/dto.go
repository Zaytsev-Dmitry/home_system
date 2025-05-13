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

type CreateAccountRequest struct {
	AccountType      *CreateAccountRequestAccountType `json:"accountType,omitempty"`
	Email            *string                          `json:"email,omitempty"`
	FirstName        *string                          `json:"firstName,omitempty"`
	LastName         *string                          `json:"lastName,omitempty"`
	TelegramId       *int64                           `json:"telegramId,omitempty"`
	TelegramUserName *string                          `json:"telegramUserName,omitempty"`
	Username         *string                          `json:"username,omitempty"`
}

type AccountResponse struct {
	Email            *string `json:"email,omitempty"`
	FirstName        *string `json:"firstName,omitempty"`
	Id               *int64  `json:"id,omitempty"`
	KeycloakId       *string `json:"keycloakId,omitempty"`
	LastName         *string `json:"lastName,omitempty"`
	TelegramId       *int64  `json:"telegramId,omitempty"`
	TelegramUserName *string `json:"telegramUserName,omitempty"`
	Username         *string `json:"username,omitempty"`
}

type ProfileResponse struct {
	AccountId *int64  `json:"accountId,omitempty"`
	Id        *int64  `json:"id,omitempty"`
	Role      *string `json:"role,omitempty"`
	Username  *string `json:"username,omitempty"`
}
