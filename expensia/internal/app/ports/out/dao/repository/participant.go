package repository

type ParticipantRepository interface {
	GetIdByTgUserId(userId int64) (int64, error)
}
