package noteDao

import (
	"fmt"
	noteDomain "noteBackendApp/internal/domain"
)

var notes = make(map[string]noteDomain.NoteEntity)

func Save(toSave noteDomain.NoteEntity) noteDomain.NoteEntity {
	if _, found := notes[toSave.Id]; !found {
		notes[toSave.Id] = toSave
	} else {
		return notes[toSave.Id]
	}
	return toSave
}

// TODO добавить логирование
func GetById(id string) (noteDomain.NoteEntity, error) {
	if obj, found := notes[id]; found {
		return obj, nil
	} else {
		errorFormatStr := "объект с id: %s не найден"
		err := fmt.Errorf(errorFormatStr, id)
		return noteDomain.NoteEntity{}, err
	}
}
