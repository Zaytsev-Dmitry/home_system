package delegate

import (
	"noteBackendApp/internal/app/domain"
	"noteBackendApp/internal/app/ports/out/dao"
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

func (n *NoteDelegate) Save(toSave domain.Note) domain.Note {
	return n.saveUCase.Save(toSave)
}

func (g *NoteDelegate) GetNoteByTgId(id int64) []domain.Note {
	return g.getUCase.GetNoteByTgId(id)
}

func (d *NoteDelegate) DeleteNoteByTgId(tgId int64) {

}
