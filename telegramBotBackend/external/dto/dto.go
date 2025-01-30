package dto

type CreateAccountRequestAccountType string

type AccountResponse struct {
	Email     *string `json:"email,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
	Login     *string `json:"login,omitempty"`
}

type CreateAccountRequest struct {
	AccountType *CreateAccountRequestAccountType `json:"accountType,omitempty"`
	Email       *string                          `json:"email,omitempty"`
	FirstName   *string                          `json:"firstName,omitempty"`
	LastName    *string                          `json:"lastName,omitempty"`
	Login       *string                          `json:"login,omitempty"`
	Password    *string                          `json:"password,omitempty"`
	TelegramId  *int64                           `json:"telegramId,omitempty"`
}

type AccountDTO struct {
	ID        uint64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Login     string `json:"login"`
	Email     string `json:"email"`
	Type      string `json:"type"`
}
