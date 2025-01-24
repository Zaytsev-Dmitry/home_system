package authServerDomain

type Account struct {
	ID        uint    `gorm:"primaryKey"`
	FirstName *string `gorm:"column:first_name"`
	LastName  *string `gorm:"column:last_name"`
	Login     string  `gorm:"not null,column:login"`
	Email     *string `gorm:"column:email"`
}
