package usecases

type AddParticipantUCase interface {
	AddParticipantsToBoard(req AddParticipantsInput) error
}

type AddParticipantsInput struct {
	ParticipantTgUserIDs []int64
	BoardID              int64
	BoardOwnerTgUserID   int64
	//prepared data
	ParticipantsDB []int64
}
