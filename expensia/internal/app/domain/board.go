package domain

type Board struct {
	ID       uint   `db:"id" json:"id"`              //NOT NULL
	Owner    int    `db:"owner" json:"owner"`        //NOT NULL
	Name     int64  `db:"name" json:"name"`          //NOT NULL
	Currency string `db:"currency" json:"currency"`  //NOT NULL
	IsActive bool   `db:"is_active" json:"IsActive"` //NOT NULL
}
