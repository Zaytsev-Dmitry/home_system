package noteDaoPorts

import (
	"fmt"
	noteDomain "noteBackendApp/internal/domain"
)

type InMemoryNoteRepository struct {
	Notes map[int]noteDomain.TelegramAccount
}

func NewInMemoryNoteRepository() *InMemoryNoteRepository {
	return &InMemoryNoteRepository{
		Notes: make(map[int]noteDomain.TelegramAccount),
	}
}

func (db *InMemoryNoteRepository) Save(toSave noteDomain.TelegramAccount) noteDomain.TelegramAccount {
	notes := db.Notes
	if _, found := db.Notes[toSave.AccountId]; !found {
		notes[toSave.AccountId] = toSave
	} else {
		return notes[toSave.AccountId]
	}
	return toSave
}

// TODO добавить логирование
func (db *InMemoryNoteRepository) GetNoteByAccountId(id int) (noteDomain.TelegramAccount, error) {
	notes := db.Notes
	if obj, found := notes[id]; found {
		return obj, nil
	} else {
		errorFormatStr := "объект с id: %s не найден"
		err := fmt.Errorf(errorFormatStr, id)
		return noteDomain.TelegramAccount{}, err
	}
}

// TODO добавить логирование
func (db *InMemoryNoteRepository) DeleteNoteByAccountId(id int) {
	notes := db.Notes
	if obj, found := notes[id]; found {
		delete(notes, obj.AccountId)
	}
}
