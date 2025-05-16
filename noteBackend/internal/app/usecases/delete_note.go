package usecases

type DeleteNoteUCase interface {
	DeleteNoteByTgId(tgId int64) error
}
