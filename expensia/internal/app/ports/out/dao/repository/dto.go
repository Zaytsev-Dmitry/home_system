package repository

type CreateBoardUCaseIn struct {
	TgUserId int64
	OwnerId  int64
	Name     string
	Currency string
}

type AddParticipantsInput struct {
	ParticipantTgUserIDs []int64
	BoardID              int64
	BoardOwnerTgUserID   int64
	//prepared data
	ParticipantsDB []int64
}
