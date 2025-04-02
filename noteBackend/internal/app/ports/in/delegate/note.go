package delegate

import (
	"noteBackendApp/internal/app/usecases"
)

type NoteDelegate struct {
	saveUCase   *usecases.SaveNoteUCase
	getUCase    *usecases.GetNoteUCase
	deleteUCase *usecases.DeleteNoteUCase
}
