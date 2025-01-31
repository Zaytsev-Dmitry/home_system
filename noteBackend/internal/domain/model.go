package noteDomain

type TelegramAccount struct {
	ID          uint    `gorm:"primaryKey"`
	AccountId   int     `gorm:"column:account_id"`
	Name        string  `gorm:"column:name"`
	Link        *string `gorm:"column:link"`
	Description *string `gorm:"column:description"`
}

type Note struct {
	ID          uint    `db:"id"`
	AccountId   int     `db:"account_id"`
	Name        string  `db:"name"`
	Link        *string `db:"link"`
	Description *string `db:"description"`
}
