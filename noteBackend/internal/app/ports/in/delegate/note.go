package delegate

import (
	"noteBackendApp/internal/app/domain"
	"noteBackendApp/internal/app/ports/out/dao"
	"noteBackendApp/internal/app/ports/out/dao/dto"
	"noteBackendApp/internal/app/services"
	"noteBackendApp/internal/app/usecases"
)

type NoteDelegate struct {
	saveUCase   usecases.SaveNoteUCase
	getUCase    usecases.GetNoteUCase
	deleteUCase usecases.DeleteNoteUCase
}

func Create(dao *dao.NoteDao) *NoteDelegate {
	return &NoteDelegate{
		saveUCase:   &services.SaveNoteUCaseImpl{NoteRepo: dao.NoteRepo},
		getUCase:    &services.GetNoteUCaseImpl{NoteRepo: dao.NoteRepo},
		deleteUCase: &services.DeleteNoteUCaseImpl{NoteRepo: dao.NoteRepo},
	}
}

func (d *NoteDelegate) Save(dto dto.CreateNoteDto) (*domain.Note, error) {
	return d.saveUCase.Save(dto)
}

func (d *NoteDelegate) GetNoteByTgId(id int64) ([]*domain.Note, error) {
	return d.getUCase.GetNoteByTgId(id)
}

func (d *NoteDelegate) DeleteNoteByTgId(tgId int64) error {
	return d.deleteUCase.DeleteNoteByTgId(tgId)
}
