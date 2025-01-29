package noteDaoInterface

import (
	noteDomain "noteBackendApp/internal/domain"
)

type NoteDao interface {
	Save(entity noteDomain.TelegramAccount) noteDomain.TelegramAccount
	DeleteNotesByAccountId(accountId int)
	GetNotesByAccountId(accountId int) []noteDomain.TelegramAccount
	ExistByName(name string) bool
	GetByName(name string) noteDomain.TelegramAccount
}
