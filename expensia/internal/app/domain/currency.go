package domain

type Currency struct {
	Id       int    `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`          //unique not null
	IsActive bool   `db:"is_active" json:"isActive"` //NOT NULL
}
