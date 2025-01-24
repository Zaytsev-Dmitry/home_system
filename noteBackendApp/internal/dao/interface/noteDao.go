package noteDaoInterface

import noteDomain "noteBackendApp/internal/domain"

type NoteDao interface {
	Save(entity noteDomain.TelegramAccount) noteDomain.TelegramAccount
}
