package noteDomain

type TelegramAccount struct {
	ID          uint    `gorm:"primaryKey"`
	AccountId   int     `gorm:"column:account_id"`
	Name        string  `gorm:"column:name"`
	Link        *string `gorm:"column:link"`
	Description *string `gorm:"column:description"`
}
