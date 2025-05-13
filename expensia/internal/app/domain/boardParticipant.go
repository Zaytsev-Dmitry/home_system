package domain

import "time"

type BoardParticipant struct {
	BoardId       int       `db:"board_id" json:"boardId"`             //NOT NULL
	ParticipantId int       `db:"participant_id" json:"participantId"` //NOT NULL
	JoinedAtUtc   time.Time `db:"joined_at_utc" json:"joinedAtUtc"`    //NOT NULL
}
