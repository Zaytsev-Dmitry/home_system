package domain

type ExpenseShare struct {
	ExpenseId     int64   `db:"expense_id" json:"expenseId"`
	ParticipantId int64   `db:"participant_id" json:"participantId"`
	ShareAmount   float64 `db:"share_amount" json:"shareAmount"`
}
