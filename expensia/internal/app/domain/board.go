package domain

type Board struct {
	ID       uint   `db:"id" json:"id"`              //NOT NULL
	OwnerId  int64  `db:"owner_id" json:"ownerId"`   //NOT NULL FK
	Name     string `db:"name" json:"name"`          //NOT NULL
	Currency string `db:"currency" json:"currency"`  //NOT NULL
	IsActive bool   `db:"is_active" json:"IsActive"` //NOT NULL
}
