package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	openapi "noteBackendApp/api/http"
	"noteBackendApp/internal/app/domain"
	"noteBackendApp/internal/app/ports/in/delegate"
	noteInterface "noteBackendApp/internal/app/ports/out/dao"
	notePresenter "noteBackendApp/internal/infrastructure/transport/http/presenter"
	"noteBackendApp/pkg/marshalling"
)

type NoteController struct {
	delegate  *delegate.NoteDelegate
	presenter notePresenter.Presenter
}

func (cntr *NoteController) SaveNote(c *gin.Context) {
	var req openapi.CreateNoteRequest
	if err := marshalling.HandleMarshalling(c, &req); err != nil {
		return
	}
	marshalling.HandleResponse(c, func() (*domain.Note, error) {
		return cntr.delegate.Save(cntr.presenter.PresentToReq(req))
	}, cntr.presenter.PresentToResp)
}

func (cntr *NoteController) DeleteNotesByTgId(c *gin.Context, tgId int64) {
	marshalling.ExecuteAndRespondStatus(
		http.StatusNoContent, c,
		func() error {
			return cntr.delegate.DeleteNoteByTgId(tgId)
		},
	)
}

func (cntr *NoteController) GetNotesByTgId(c *gin.Context, tgId int64) {
	marshalling.HandleResponse(c, func() ([]*domain.Note, error) {
		return cntr.delegate.GetNoteByTgId(tgId)
	}, cntr.presenter.PresentToListResp)
}

func Create(dao *noteInterface.NoteDao) *NoteController {
	return &NoteController{
		delegate:  delegate.Create(dao),
		presenter: notePresenter.Presenter{},
	}
}
