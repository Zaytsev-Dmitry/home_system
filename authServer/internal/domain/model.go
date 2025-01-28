package authServerDomain

type Account struct {
	ID        uint64  `db:"id"`
	FirstName *string `db:"first_name"`
	LastName  *string `db:"last_name"`
	Login     string  `db:"login"`
	Email     *string `db:"email"`
}
