package controller

import (
	"github.com/gin-gonic/gin"
	"noteBackendApp/internal/app/ports/in/delegate"
	noteInterface "noteBackendApp/internal/app/ports/out/dao"
	notePresenter "noteBackendApp/internal/infrastructure/transport/http/presenter"
)

type NoteController struct {
	delegate  *delegate.NoteDelegate
	presenter notePresenter.Presenter
}

func (controller *NoteController) SaveNote(c *gin.Context) {

}

func (controller *NoteController) DeleteNotesByTgId(context *gin.Context, tgId int64) {
}

func (controller *NoteController) GetNotesByTgId(context *gin.Context, tgId int64) {
}

func Create(dao *noteInterface.NoteDao) *NoteController {
	return &NoteController{
		delegate:  delegate.Create(dao),
		presenter: notePresenter.Presenter{},
	}
}
