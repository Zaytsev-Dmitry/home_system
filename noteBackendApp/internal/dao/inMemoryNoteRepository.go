package noteDao

import (
	"fmt"
	"maps"
	noteDomain "noteBackendApp/internal/domain"
	"slices"
)

type InMemoryNoteRepository struct {
	Notes map[string]noteDomain.NoteEntity
}

func NewInMemoryNoteRepository() *InMemoryNoteRepository {
	return &InMemoryNoteRepository{
		Notes: make(map[string]noteDomain.NoteEntity),
	}
}

func (db *InMemoryNoteRepository) Save(toSave noteDomain.NoteEntity) noteDomain.NoteEntity {
	notes := db.Notes
	if _, found := db.Notes[toSave.Id]; !found {
		notes[toSave.Id] = toSave
	} else {
		return notes[toSave.Id]
	}
	return toSave
}

// TODO добавить логирование
func (db *InMemoryNoteRepository) GetById(id string) (noteDomain.NoteEntity, error) {
	notes := db.Notes
	if obj, found := notes[id]; found {
		return obj, nil
	} else {
		errorFormatStr := "объект с id: %s не найден"
		err := fmt.Errorf(errorFormatStr, id)
		return noteDomain.NoteEntity{}, err
	}
}

// TODO добавить логирование
func (db *InMemoryNoteRepository) DeleteById(id string) {
	notes := db.Notes
	if obj, found := notes[id]; found {
		delete(notes, obj.Id)
	}
}

// TODO добавить логирование
func (db *InMemoryNoteRepository) GetAll() []noteDomain.NoteEntity {
	return slices.Collect(maps.Values(db.Notes))
}
