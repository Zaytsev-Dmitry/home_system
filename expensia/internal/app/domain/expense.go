package domain

import "time"

type Expense struct {
	Id            int       `db:"id" json:"id"`
	Title         string    `db:"title" json:"title"`                  //NOT NULL
	ParticipantId int64     `db:"participant_id" json:"participantId"` //NOT NULL
	BoardId       int64     `db:"board_id" json:"boardId"`             //NOT NULL
	Amount        float64   `db:"amount" json:"amount"`                //NOT NULL
	CreatedAtUtc  time.Time `db:"created_at_utc" json:"createdAtUtc"`  //NOT NULL
	CategoryId    int64     `db:"category_id" json:"categoryId"`       //NOT NULL
}
