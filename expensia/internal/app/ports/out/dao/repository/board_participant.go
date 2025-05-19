package repository

type BoardParticipantRepository interface {
	AddParticipantsToBoard(boardId int64, participantIds []int64) error
}
